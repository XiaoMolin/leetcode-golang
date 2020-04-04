// 两数之和 II - 输入有序数组 project main.go
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
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			result := []int{i + 1, j + 1}
			return result
		}
	}
	return []int{}
}
