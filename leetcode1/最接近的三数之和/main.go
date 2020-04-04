// 最接近的三数之和 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
	var nums = []int{21, 54, 23, 50, 72, 21, 53, 83, 54, 99, 43}
	target := 3
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		L := i + 1
		R := len(nums) - 1
		for {
			if L >= R {
				break
			}
			sum := nums[i] + nums[L] + nums[R]
			if sum > target {
				R--
			} else if sum < target {
				L++
			} else {
				return target
			}

			if distanceInt(sum, target) < distanceInt(res, target) {
				res = sum
			}
		}
	}

	return res
}

func distanceInt(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}
