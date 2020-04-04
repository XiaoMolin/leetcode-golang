// 买卖股票的最佳时机含手续费 project main.go
package main

import (
	"fmt"
)

func main() {
	//prices := []int{1, 3, 2, 8, 4, 9}
	//fee := 2
	prices := []int{1, 3, 7, 5, 10, 3} //0，4，2，7，0
	fee := 3

	fmt.Println(maxProfit(prices, fee))
}

/*因为添加了手续费，所以贪心算法失败*/
/*
执行用时 :112 ms, 在所有 Go 提交中击败了86.96%的用户
内存消耗 :7.4 MB, 在所有 Go 提交中击败了74.42%的用户

动态规划
*/

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-fee+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	//fmt.Println(dp)
	return dp[len(prices)-1][0]
}

/*
优化后
执行用时 :116 ms, 在所有 Go 提交中击败了67.39%的用户
内存消耗 :7.2 MB, 在所有 Go 提交中击败了90.70%的用户
*/
func maxProfit1(prices []int, fee int) int {
	res := 0
	hold := -prices[0]
	for i := 1; i < len(prices); i++ {
		res = max(res, hold+prices[i]-fee)
		hold = max(hold, res-prices[i])
	}
	return res
}
func max(l int, r int) int {
	if l > r {
		return l
	} else {
		return r
	}
}
