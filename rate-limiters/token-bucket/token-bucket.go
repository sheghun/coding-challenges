package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter manages token buckets for multiple clients (like Cloudflare)
type RateLimiter struct {
	buckets         map[string]*TokenBucket
	mutex           sync.RWMutex
	capacity        int           // max tokens per bucket
	refillRate      float64       // tokens per second
	cleanupInterval time.Duration // how often to cleanup inactive buckets
	lastCleanup     time.Time
}

// TokenBucket represents a single client's token bucket
type TokenBucket struct {
	tokens         float64   // current token count (float for precision)
	lastRefillTime time.Time // last time tokens were refilled
	capacity       int       // maximum tokens
	refillRate     float64   // tokens per second
}

// NewRateLimiter creates a new distributed rate limiter
func NewRateLimiter(capacity int, refillRate float64) *RateLimiter {
	return &RateLimiter{
		buckets:         make(map[string]*TokenBucket),
		capacity:        capacity,
		refillRate:      refillRate,
		cleanupInterval: 10 * time.Minute,
		lastCleanup:     time.Now(),
	}
}

// Allow checks if a request from clientID should be allowed
func (rl *RateLimiter) Allow(clientID string) bool {
	return rl.AllowN(clientID, 1)
}

// AllowN checks if N tokens are available for clientID
func (rl *RateLimiter) AllowN(clientID string, tokens int) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	// Get or create bucket for this client
	bucket := rl.getBucket(clientID)

	// Refill tokens based on time elapsed
	rl.refillBucket(bucket)

	// Check if we have enough tokens
	if bucket.tokens >= float64(tokens) {
		bucket.tokens -= float64(tokens)
		return true
	}

	return false
}

// getBucket gets existing bucket or creates new one for client
func (rl *RateLimiter) getBucket(clientID string) *TokenBucket {
	bucket, exists := rl.buckets[clientID]
	if !exists {
		bucket = &TokenBucket{
			tokens:         float64(rl.capacity), // start with full bucket
			lastRefillTime: time.Now(),
			capacity:       rl.capacity,
			refillRate:     rl.refillRate,
		}
		rl.buckets[clientID] = bucket
	}
	return bucket
}

// refillBucket adds tokens based on time elapsed since last refill
func (rl *RateLimiter) refillBucket(bucket *TokenBucket) {
	now := time.Now()
	elapsed := now.Sub(bucket.lastRefillTime).Seconds()

	// Calculate tokens to add: elapsed_seconds * tokens_per_second
	tokensToAdd := elapsed * bucket.refillRate

	// Add tokens but don't exceed capacity
	bucket.tokens = min(bucket.tokens+tokensToAdd, float64(bucket.capacity))
	bucket.lastRefillTime = now
}

// GetStats returns current stats for a client (useful for monitoring)
func (rl *RateLimiter) GetStats(clientID string) (currentTokens float64, exists bool) {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()

	bucket, exists := rl.buckets[clientID]
	if !exists {
		return 0, false
	}

	// Create a copy to refill without modifying original
	tempBucket := *bucket
	rl.refillBucket(&tempBucket)

	return tempBucket.tokens, true
}

// Cleanup removes inactive clients to prevent memory leaks
func (rl *RateLimiter) Cleanup() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	if now.Sub(rl.lastCleanup) < rl.cleanupInterval {
		return
	}

	// Remove buckets that haven't been accessed in a while
	cutoff := now.Add(-rl.cleanupInterval)
	for clientID, bucket := range rl.buckets {
		if bucket.lastRefillTime.Before(cutoff) {
			delete(rl.buckets, clientID)
		}
	}

	rl.lastCleanup = now
}

// ActiveClients returns number of active clients
func (rl *RateLimiter) ActiveClients() int {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	return len(rl.buckets)
}

// min helper function (Go 1.21+ has this built-in)
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Example usage demonstrating Cloudflare-style rate limiting
func main() {
	// Create rate limiter: 10 requests per second, burst of 100
	limiter := NewRateLimiter(100, 10.0)

	// Simulate requests from different clients
	clients := []string{"192.168.1.1", "10.0.0.1", "203.0.113.1"}

	fmt.Println("=== Rate Limiter Demo ===")

	for i := 0; i < 15; i++ {
		for _, client := range clients {
			allowed := limiter.Allow(client)
			tokens, _ := limiter.GetStats(client)

			status := "ALLOWED"
			if !allowed {
				status = "BLOCKED"
			}

			fmt.Printf("Client %s: %s (%.2f tokens remaining)\n",
				client, status, tokens)
		}

		fmt.Printf("Active clients: %d\n", limiter.ActiveClients())
		fmt.Println("---")

		// Small delay to show token refill
		time.Sleep(100 * time.Millisecond)

		// Periodic cleanup
		limiter.Cleanup()
	}
}
