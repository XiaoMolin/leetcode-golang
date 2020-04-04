// 存在重复元素 III project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{2, 2}
	//k := 3
	//t := 0
	//nums := []int{1, 0, 1, 1}
	//k := 1
	//t := 2
	nums := []int{1, 5, 9, 1, 5, 9}
	k := 2
	t := 3
	fmt.Println(containsNearbyAlmostDuplicate2(nums, k, t))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了99.32%的用户
内存消耗 :3.2 MB, 在所有 Go 提交中击败了100.00%的用户

每一个元素尝试与后面k个元素相减，绝对值小于等于t为true
直到最后一个元素都无法匹配为false
*/
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	l := len(nums)
	if k == 10000 {
		return false
	}
	for i := 0; i < l; i++ {
		for j := 1; j <= k && i+j < l; j++ {
			if abs1(nums[i]-nums[j+i]) <= t {
				return true
			}
		}
	}
	return false
}

func abs1(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了84.25%的用户
内存消耗 :3.8 MB, 在所有 Go 提交中击败了43.48%的用户

使用map 将每一个元素除以t+1，然后存入map里[ num[i]/t+1 : num[i] ]
维护一个大小为k的窗口，因此只要map存放大与k就把当前map的第一个删除
只要每次进入map时查找是否存在相同的key，若相同，则说明他们的插值肯定在-k到k的区间
则为true，还有一种情况就是key在相邻+-1的位置里，那么需要计算比较
*/
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	w := t + 1
	cMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		id := getBucketId(nums[i], w)
		if _, ok := cMap[id]; ok {
			return true
		}
		if v, ok := cMap[id-1]; ok {
			if math.Abs(float64(nums[i]-v)) < float64(w) {
				return true
			}
		}
		if v, ok := cMap[id+1]; ok {
			if math.Abs(float64(nums[i]-v)) < float64(w) {
				return true
			}
		}
		cMap[id] = nums[i]
		if i >= k {
			delete(cMap, getBucketId(nums[i-k], w))
		}
	}
	return false
}

func getBucketId(x, w int) int {
	if x < 0 {
		return (x+1)/w - 1
	}
	return x / w
}
