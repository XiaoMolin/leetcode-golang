// 买卖股票的最佳时机 II project main.go
package main

import (
	"fmt"
)

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4} //7
	//prices := []int{1, 2, 3, 4, 5} //4
	//prices := []int{7, 6, 4, 3, 1} //0
	prices := []int{2, 4, 1} //2
	//prices := []int{3, 2, 6, 5, 0, 3} //7
	fmt.Println(maxProfit(prices))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了96.64%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户

low 和high 记录每次买入卖出的最低值和最高值
遍历，寻找每一次峰值之前的最小值买入，
即low为一次峰值中的最小值，high为峰顶
利润 res+=high-low
返回res即可
*/
func maxProfit_0(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	high := 0
	low := prices[0]
	res := 0
	for i := 0; i < l; i++ {
		if prices[i] > high && prices[i] > low {
			high = prices[i]
		} else if low < high {
			res = res + high - low
			high = 0
			low = prices[i]
		}
		if low > prices[i] {
			low = prices[i]
		}
		if i == l-1 && low < high {
			res = res + high - low
		}
	}
	return res
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了96.64%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了59.61%的用户
贪心算法，一次遍历，只要今天价格小于明天价格就在今天买入然后明天卖出，时间复杂度O(n)
*/

func maxProfit_1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

/*
动态规划
执行用时 :4 ms, 在所有 Go 提交中击败了96.79%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了59.61%的用户
*/
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	for i := 1; i < l; i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
