package main

import (
	"fmt"
	"sync"
)

func Job() chan int {
	ch := make(chan int, 100)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func Worker(wg *sync.WaitGroup) {
	for val := range Job() {
		wg.Add(1)
		go Process(val)
		wg.Done()
	}
}

func Process(i int) {
	for j := 0; j < i; j++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	Worker(&wg)
	wg.Wait()
	fmt.Println("done")
}
