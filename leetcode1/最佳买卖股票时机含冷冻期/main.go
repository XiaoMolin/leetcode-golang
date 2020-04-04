// 最佳买卖股票时机含冷冻期 project main.go
package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.3 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	dp_pre := 0
	for i := 1; i < l; i++ {
		temp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp_pre-prices[i])
		dp_pre = temp
	}
	return dp0
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
