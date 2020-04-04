// 数组的相对排序 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了96.61%的用户

遍历一遍arr2，再每一个arr2元素期间遍历arr1只要相等就按顺序放到arr1里
arr1[lastIdx], arr1[j] = arr1[j], arr1[lastIdx]
后面的排个序就行sort.Ints(arr1[lastIdx:])
时间复杂度O(N*N)
空间复杂度O(N)
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	lastIdx := 0
	for i := 0; i < len(arr2); i++ {
		for j := lastIdx; j < len(arr1); j++ {
			if arr1[j] == arr2[i] {
				arr1[lastIdx], arr1[j] = arr1[j], arr1[lastIdx]
				lastIdx++
			}
		}
	}
	sort.Ints(arr1[lastIdx:])
	return arr1
}
