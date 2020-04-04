// 使用最小花费爬楼梯 project main.go
/*
数组的每个索引做为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。
每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。
您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:
输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。

 示例 2:
输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。

注意：
cost 的长度将会在 [2, 1000]。
每一个 cost[i] 将会是一个Integer类型，范围为 [0, 999]。

*/
package main

import (
	"fmt"
)

func main() {
	cost := []int{1, 2, 2, 2, 332, 252, 21, 23, 4, 5, 65, 34, 3, 4}
	fmt.Println(minCostClimbingStairs(cost))
}

/*
动态规划
cost数组前面补一个零
0 1 2 2 ....
对于阶梯i，由于一次可以上一阶或两阶，因此必定需要从i - 1阶或i - 2阶上来(0, 1阶除外，这是边界)

dp[2]=dp[0]+cost[2]
dp[2]=dp[1]+cost[2]
选取最低的那个
dp[i]=cost[i]当i=0，1时
dp[i]=min(dp[i-2],dp[i-1])+cost当i>1时
可能会有人迷惑 可以选择从索引为 0 或 1 的元素作为初始阶梯
这一点，其实可以当做从平地作为起点，登一步到阶梯0，两步到阶梯1

执行用时 :8 ms, 在所有 Go 提交中击败了14.48%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了11.11%的用户
*/
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = min(cost[0]+cost[1], cost[1])
	l := len(cost)
	for i := 2; i < l; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[l-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

/*
优化dp数组,维护a,b两个数
执行用时 :4 ms, 在所有 Go 提交中击败了95.02%的用户
内存消耗 :2.9 MB, 在所有 Go 提交中击败了51.85%的用户
*/
func minCostClimbingStairs2(cost []int) int {
	cost = append(cost, 0)
	a := cost[0]
	b := min(cost[0]+cost[1], cost[1])
	l := len(cost)
	for i := 2; i < l; i++ {
		c := min(a, b) + cost[i]
		a = b
		b = c
	}
	return b
}
