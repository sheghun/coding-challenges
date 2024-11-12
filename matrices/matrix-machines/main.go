package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type City struct {
	value   int32
	machine bool
	roads   []*Road
}

type Road struct {
	time int32
	dest *City
}

/*
 * Complete the 'minTime' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY roads
 *  2. INTEGER_ARRAY machines
 */

var cities = make(map[int32]*City)

func minTime(roads [][]int32, machines []int32) int32 {
	// Write your code here
	for _, r := range roads {
		s, d, t := r[0], r[1], r[2]

		start, ok := cities[s]
		if !ok {
			start = &City{
				value:   s,
				machine: slices.Contains(machines, s),
				roads:   []*Road{},
			}
			cities[s] = start
		}

		dest, ok := cities[d]
		if !ok {
			dest = &City{
				value:   d,
				machine: slices.Contains(machines, d),
				roads:   []*Road{},
			}
			cities[d] = dest
		}

		start.roads = append(start.roads, &Road{
			time: t,
			dest: dest,
		})
	}

	timeTaken := []int32{}

	for _, city := range cities {
		for i := 0; i < len(city.roads); i++ {
			var time int32
			for j := 0; j < len(city.roads); j++ {
				road := city.roads[j]
				if !road.dest.machine {
					continue
				}
				time += road.time
			}
			timeTaken = append(timeTaken, time)
		}
	}

	fmt.Println(timeTaken)
	fmt.Printf("%#v", cities)
	return timeTaken[0]
}

func main() {
	os.Setenv("OUTPUT_PATH", "./output.txt")
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var roads [][]int32
	for i := 0; i < int(n)-1; i++ {
		roadsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var roadsRow []int32
		for _, roadsRowItem := range roadsRowTemp {
			roadsItemTemp, err := strconv.ParseInt(roadsRowItem, 10, 64)
			checkError(err)
			roadsItem := int32(roadsItemTemp)
			roadsRow = append(roadsRow, roadsItem)
		}

		if len(roadsRow) != 3 {
			panic("Bad input")
		}

		roads = append(roads, roadsRow)
	}

	var machines []int32

	for i := 0; i < int(k); i++ {
		machinesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		machinesItem := int32(machinesItemTemp)
		machines = append(machines, machinesItem)
	}

	result := minTime(roads, machines)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
