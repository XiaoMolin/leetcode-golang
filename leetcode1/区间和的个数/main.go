// 区间和的个数 project main.go 不是很多懂
package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 5, -1}
	lower := -2
	upper := 2
	fmt.Println(countRangeSum(nums, lower, upper))
}

/*
执行用时 :100 ms, 在所有 Go 提交中击败了33.33%的用户
内存消耗 :3.8 MB, 在所有 Go 提交中击败了93.33%的用户

暴力法
*/
func countRangeSum1(nums []int, lower int, upper int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < l; j++ {
			sum += nums[j]
			if sum <= upper && sum >= lower {
				res++
			}
		}
	}
	return res
}

/*
思路二：
如果满足条件的范围的右端点是当前节点，那么有多少区间是满足条件的呢？
设pre_sum[i]表示从开始位置到当前位置的前缀和，我们可以从0位置开始，
到i-1这个范围中，通过pre_sum[i]减去前面的前缀和，就得到了一个区间的和

如果区间和是从i到j，j是当前需要判断的区间右节点
lower <= pre_sum[j] - pre_sum[i] <= upper
变换可得:
pre_sum[j] - upper <= pre_sum[i] <= pre_sum[j] - lower

*/

/*
执行用时 :116 ms, 在所有 Go 提交中击败了33.33%的用户
内存消耗 :5.7 MB, 在所有 Go 提交中击败了66.67%的用户

*/
func countRangeSum2(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	S := []int{0}
	for i := 0; i < len(nums); i++ {
		S = append(S, S[i]+nums[i])
	}
	fmt.Println(S)
	for i := 1; i < len(S); i++ {
		for j := 0; j < i; j++ {
			if S[i]-S[j] <= upper && S[i]-S[j] >= lower {
				res++
			}
		}
	}
	return res
}

func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	assist := []int{}
	for i := 0; i < len(nums); i++ {
		assist = append(assist, 0)
	}
	S := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		S = append(S, S[i-1]+nums[i])
	}
	fmt.Println(S)
	return merge(S, 0, len(nums)-1, assist, lower, upper)
}

/*
执行用时 :28 ms, 在所有 Go 提交中击败了55.56%的用户
内存消耗 :6.1 MB, 在所有 Go 提交中击败了40.00%的用户

归并排序


*/
func merge(nums []int, l int, r int, assist []int, low int, up int) int {
	if l >= r {
		if l == r && nums[l] >= low && nums[l] <= up {
			return 1
		}
		return 0
	}
	cnt := 0
	mid := (l + r) >> 1
	cnt += merge(nums, l, mid, assist, low, up)
	cnt += merge(nums, mid+1, r, assist, low, up)

	left := l
	lower, upper := mid+1, mid+1
	for left <= mid {
		for lower <= r && nums[lower]-nums[left] < low {
			lower++
		}
		for upper <= r && nums[upper]-nums[left] <= up {
			upper++
		}
		cnt += (upper - lower)
		left++
	}
	fmt.Println(nums, l, r, cnt)
	left = l
	right := mid + 1
	pos := l
	//合并
	for left <= mid || right <= r {
		if left > mid {
			assist[pos] = nums[right]
			right++
		} else if right > r {
			assist[pos] = nums[left]
			left++
		} else if nums[left] <= nums[right] {
			assist[pos] = nums[left]
			left++
		} else {
			assist[pos] = nums[right]
			right++
		}
		pos++
	}
	for i := l; i <= r; i++ {
		nums[i] = assist[i]
	}
	return cnt
}
