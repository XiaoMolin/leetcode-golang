// 买卖股票的最佳时机 project main.go
package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了96.77%的用户
内存消耗 :3.6 MB, 在所有 Go 提交中击败了5.02%的用户
*/
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}

func max(l int, r int) int {
	if l > r {
		return l
	} else {
		return r
	}
}

/*
优化后
执行用时 :4 ms, 在所有 Go 提交中击败了96.77%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, -prices[i])
	}
	return dp0
}
