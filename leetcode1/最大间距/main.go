// 最大间距 project main.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{32, 63, 91, 12, 1, 3, 999}
	fmt.Println(maximumGap(nums))
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了44.26%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了97.06%的用户

比较排序
时间复杂度：O(n\log n)O(nlogn)。排序的复杂度是 O(n\log n)O(nlogn)，
遍历的复杂度是 O(n)O(n)，总复杂度是 O(n \log n)O(nlogn)。

空间复杂度：除去输入数组之外，不需要额外空间（因为大多数都是原地排序）。

*/
func maximumGap1(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	for i := 0; i < l-1; i++ {
		if nums[i+1]-nums[i] > res {
			res = nums[i+1] - nums[i]
		}
	}
	return res
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了42.62%的用户
内存消耗 :5.9 MB, 在所有 Go 提交中击败了5.88%的用户

基数排序
*/
func maximumGap2(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	res := 0
	barrel := make([][]int, 10)
	max := nums[0]
	//找最大数确定位数
	for i := 1; i < l; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	exp := 1 //一位一位的进行
	for max > 0 {
		for i := 0; i < 10; i++ { //将之前的元素清空
			barrel[i] = []int{}
		}
		for i := 0; i < l; i++ { //将数字放入对应的位置
			barrel[(nums[i]/exp)%10] = append(barrel[(nums[i]/exp)%10], nums[i])
		}
		index := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < len(barrel[i]); j++ {
				nums[index] = barrel[i][j]
				index++
			}
		}
		max /= 10
		exp *= 10
	}
	fmt.Println(nums)
	for i := 0; i < l-1; i++ {
		if nums[i+1]-nums[i] > res {
			res = nums[i+1] - nums[i]
		}
	}
	return res
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了42.62%的用户
内存消耗 :3.4 MB, 在所有 Go 提交中击败了61.76%的用户

桶排序
*/
func maximumGap(nums []int) int {
	// 桶排序的应用，将 nums 中的数字分配到若干个桶中，保证 maxGap 不会出现在桶中(即桶大小
	// 小于等于最小的 maxGap)，最后遍历所有桶，两个桶之间的 gap 为后一个桶中的最小元素减去
	// 前一个桶中的最大元素，并获取最大的 gap 即可。

	// 首先需要确定桶的大小和个数，由于桶的大小不大于最小的 mapGap，所以可以设置桶的大小为
	// 最小的 maxGap，也就是 (maxNum-minNum)/(length-1) 的结果向上取整，桶的个数为 n-1。
	// 对每个元素来说，它本身的值除以桶的大小就是它对应的桶编号。由于计算两个桶之间的 gap
	// 需要每个桶中的最大元素和最小元素，所以需要两个数组分别保存。

	// 然后遍历所有桶，求出最大的 gap 即可.

	if len(nums) < 2 {
		return 0
	}
	// 求最大值和最小值
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
		minNum = min(minNum, num)
	}
	// 最小的 maxGap
	minMaxGap := int(math.Ceil(float64(maxNum-minNum) / float64(len(nums)-1)))

	// 分别保存每个桶中的最大值和最小值
	maxOfEachBucket := make([]int, len(nums)-1)
	minOfEachBucket := make([]int, len(nums)-1)
	fill(maxOfEachBucket, math.MinInt32)
	fill(minOfEachBucket, math.MaxInt32)

	for _, num := range nums {
		// 跳过 minNum 和 maxNum
		if num == maxNum || num == minNum {
			continue
		}
		// 更新每个桶中的最大值和最小值
		bucketID := (num - minNum) / minMaxGap
		maxOfEachBucket[bucketID] = max(maxOfEachBucket[bucketID], num)
		minOfEachBucket[bucketID] = min(minOfEachBucket[bucketID], num)
	}

	left := minNum
	maxGap := minMaxGap
	for i := 0; i < len(nums)-1; i++ {
		// 跳过空桶
		if maxOfEachBucket[i] == math.MinInt32 || minOfEachBucket[i] == math.MaxInt32 {
			continue
		}
		maxGap = max(maxGap, minOfEachBucket[i]-left)
		left = maxOfEachBucket[i]
	}
	// maxNum 单独计算
	maxGap = max(maxGap, maxNum-left)

	return maxGap
}

func fill(nums []int, num int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = num
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
