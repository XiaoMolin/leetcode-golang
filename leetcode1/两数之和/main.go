// 两数之和 project main.go
package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	//nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	result := make([]int, 0)
	for index, num := range nums {
		_, isTrue := maps[target-num]
		if isTrue {
			result = append(result, maps[target-num])
			result = append(result, index)
			break
		}
		maps[num] = index
	}
	return result
}
