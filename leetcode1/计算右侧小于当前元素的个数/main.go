// 计算右侧小于当前元素的个数 project main.go
package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 2, 3, 6, 1}
	fmt.Println(countSmaller(nums))
}

/*
执行用时 :160 ms, 在所有 Go 提交中击败了58.33%的用户
内存消耗 :4.9 MB, 在所有 Go 提交中击败了58.06%的用户

暴力法
*/
func countSmaller1(nums []int) []int {
	res := []int{}
	l := len(nums)
	for i := 0; i < l; i++ {
		a := 0
		for j := i; j < l; j++ {
			if nums[j] < nums[i] {
				a++
			}
		}
		res = append(res, a)
	}
	return res
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了99.01%的用户
内存消耗 :4.5 MB, 在所有 Go 提交中击败了74.19%的用户

*/
// 分治+索引数组
func countSmaller(nums []int) []int {
	tmp := make([]int, len(nums))   // 临时数组
	res := make([]int, len(nums))   // 结果数组
	index := make([]int, len(nums)) // 索引数组
	for i := 0; i < len(nums); i++ {
		index[i] = i // 初始化索引数组
	}
	helper2(nums, 0, len(nums)-1, index, tmp, res)
	return res
}

func helper2(nums []int, left, right int, index, temp, res []int) {
	// terminator
	if left >= right {
		return
	}
	// prepare data
	mid := left + (right-left)>>1
	// conquer sub problems
	helper2(nums, left, mid, index, temp, res)
	helper2(nums, mid+1, right, index, temp, res)
	// resolve result
	// left-right
	merge2(nums, left, mid, right, index, temp, res)
}

// index是索引数组，tmp是额外操作用的数组
func merge2(arr []int, low, mid, high int, index, tmp, res []int) {
	i, j, k := low, mid+1, low
	for i <= mid && j <= high {
		if arr[index[i]] <= arr[index[j]] {
			tmp[k] = index[j]
			k, j = k+1, j+1
		} else {
			res[index[i]] += high - j + 1
			tmp[k] = index[i]
			k, i = k+1, i+1
		}
	}
	for i <= mid {
		tmp[k] = index[i]
		k, i = k+1, i+1
	}
	for j <= high {
		tmp[k] = index[j]
		k, j = k+1, j+1
	}
	copy(index[low:high+1], tmp[low:high+1])
}
