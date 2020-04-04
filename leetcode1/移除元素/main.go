// 移除元素 project main.go
package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了8.42%的用户
*/
func removeElement(nums []int, val int) int {
	j := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums = nums[j:]
	return j
}
