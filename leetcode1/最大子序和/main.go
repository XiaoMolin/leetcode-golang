// 最大子序和 project main.go
/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1}
	fmt.Println(maxSubArray(nums))
}

/*
-2 1 -3 4

假设只有2个数字，那么最大为dp[0],dp[1],dp[0]+dp[1]
假设有3个数字,那么最大数为dp[0],dp[1],dp[2],dp[0]+dp[1],dp[1]+dp[2],dp[1]+dp[2]+dp[3]

超时
*/

func maxSubArray1(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	max := -2147483647
	for i := l - 1; i >= 0; i-- {
		dp[i][i] = nums[i]
		if dp[i][i] > max {
			max = dp[i][i]
		}
		for j := i + 1; j < l; j++ {
			dp[i][j] = dp[i][j-1] + nums[j]
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	fmt.Println(dp)
	return max
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了97.00%的用户
内存消耗 :3.3 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func maxSubArray(nums []int) int {
	sum, b := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if b > 0 {
			b += nums[i]
		} else {
			b = nums[i]
		}
		if b > sum {
			sum = b
		}
	}
	return sum
}
