// 两个数组的交集 project main.go
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了89.64%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了35.71%的用户
*/
func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for _, value := range nums1 {
		_, ok := m[value]
		if !ok {
			m[value] = value
		}
	}
	for _, value := range nums2 {
		_, ok := m[value]
		if ok {
			delete(m, value)
			res = append(res, value)
		}
	}
	return res
}
