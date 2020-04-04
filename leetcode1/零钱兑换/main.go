// 零钱兑换 project main.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

/*
贪心+dfs
执行用时 :4 ms, 在所有 Go 提交中击败了99.72%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
1.贪心
11. 想要总硬币数最少，肯定是优先用大面值硬币，所以对 coins 按从大到小排序
12. 先丢大硬币，再丢会超过总额时，就可以递归下一层丢的是稍小面值的硬币

2.乘法对加法的加速
21. 优先丢大硬币进去尝试，也没必要一个一个丢，可以用乘法算一下最多能丢几个

k = amount / coins[c_index] 计算最大能投几个
amount - k * coins[c_index] 减去扔了 k 个硬币
count + k 加 k 个硬币

3.如果因为丢多了导致最后无法凑出总额，再回溯减少大硬币数量
最先找到的并不是最优解
31. 注意不是现实中发行的硬币，面值组合规划合理，会有奇葩情况
32. 考虑到有 [1,7,10] 这种用例，按照贪心思路 10 + 1 + 1 + 1 + 1 会比 7 + 7 更早找到
33. 所以还是需要把所有情况都递归完

4.ans 疯狂剪枝
41. 贪心虽然得不到最优解，但也不是没用的
42. 我们快速算出一个贪心的 ans 之后，虽然还会有奇葩情况，但是绝大部分普通情况就可以疯狂剪枝了

*/

var ans int

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 降序排序
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	ans = math.MaxInt32
	util(amount, 0, 0, coins)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func util(amount, count, index int, coins []int) {
	if amount == 0 {
		if count < ans {
			ans = count
		}
		return
	}

	if index == len(coins) {
		return
	}

	for i := amount / coins[index]; i >= 0 && i+count < ans; i-- {
		util(amount-i*coins[index], count+i, index+1, coins)
	}
}

/*
动态规划dp[i] = min(dp[i], dp[i-coins[j]]+1)
dp[i]表示凑成总金额为i所需的最少的硬币个数
每个dp[i]都是dp[i-coins[j]]+1得来的
执行用时 :20 ms, 在所有 Go 提交中击败了38.22%的用户
内存消耗 :6 MB, 在所有 Go 提交中击败了63.35%的用户
*/
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}
			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}
