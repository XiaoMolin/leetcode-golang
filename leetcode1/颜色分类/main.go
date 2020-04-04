// 颜色分类 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了61.20%的用户
*/
func sortColors1(nums []int) {
	sort.Ints(nums)
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户

时间复杂度O(N)
空间复杂度O(1)
*/
func sortColors2(nums []int) {
	a1, a2, a0 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			a0++
		}
		if nums[i] == 1 {
			a1++
		}
		if nums[i] == 2 {
			a2++
		}
	}
	for i := 0; i < len(nums); i++ {
		if a0 > 0 {
			nums[i] = 0
			a0--
		} else if a1 > 0 {
			nums[i] = 1
			a1--
		} else if a2 > 0 {
			nums[i] = 2
			a2--
		}
	}
}

/*
 荷兰国旗问题
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了61.20%的用户

三指针法
我们用三个指针（p0, p2 和curr）来分别追踪0的最右边界，2的最左边界和当前考虑的元素。
初始化0的最右边界：p0 = 0。在整个算法执行过程中 nums[idx < p0] = 0.
初始化2的最左边界 ：p2 = n - 1。在整个算法执行过程中 nums[idx > p2] = 2.
初始化当前考虑的元素序号 ：curr = 0.
While curr <= p2 :
若 nums[curr] = 0 ：交换第 curr个 和 第p0个 元素，并将指针都向右移。
若 nums[curr] = 2 ：交换第 curr个和第 p2个元素，并将 p2指针左移 。
若 nums[curr] = 1 ：将指针curr右移。

时间复杂度O(N)
空间复杂度O(1)
*/
func sortColors(nums []int) {
	p0, curr, p2 := 0, 0, len(nums)-1
	for curr <= p2 {
		if nums[curr] == 0 {
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[p2], nums[curr] = nums[curr], nums[p2]
			p2--
		} else {
			curr++
		}
	}
}
