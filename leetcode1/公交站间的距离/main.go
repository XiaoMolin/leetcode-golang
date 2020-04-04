// 公交站间的距离 project main.go
package main

import (
	"fmt"
)

func main() {
	var distance = []int{3, 14, 5, 2, 21, 12, 17, 24, 11, 16, 15, 4, 9}
	start := 3
	destination := 1
	fmt.Println(distanceBetweenBusStops(distance, start, destination))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var result int
	i := start
	j := destination
	result1 := 0
	result2 := 0
	if start == start {
		result = 0
	}
	if start < start {
		for x := j - i; x > 0; {
			result1 += distance[i]
			i++
			x--
		}
		i = start
		for y := len(distance) + i - j; y > 0; {
			result2 += distance[i-1]
			i--
			if i == 0 {
				i = len(distance)
			}
			y--
		}
	}
	if start > start {
		for x := len(distance) - i + j; x > 0; {
			result1 += distance[i]
			i++
			if i == len(distance) {
				i = 0
			}
			x--
		}
		i = start
		for y := i - j; y > 0; {
			result2 += distance[i-1]
			i--
			y--
		}
	}
	if result1 < result2 {
		result = result1
	} else {
		result = result2
	}
	return result
}
