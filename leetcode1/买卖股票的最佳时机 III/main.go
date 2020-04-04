// 买卖股票的最佳时机 III project main.go
package main

import (
	"fmt"
)

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了99.41%的用户
内存消耗 :4.5 MB, 在所有 Go 提交中击败了46.67%的用户
*/
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][3][2]int, l)
	for i := 0; i < l; i++ {
		for k := 2; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[l-1][2][0]
}

func max(l int, r int) int {
	if l > r {
		return l
	} else {
		return r
	}
}
