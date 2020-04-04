// 石子游戏 project main.go
/*
亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。
游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。
亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。
这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。
假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。

示例：
输入：[5,3,4,5]
输出：true
解释：
亚历克斯先开始，只能拿前 5 颗或后 5 颗石子 。
假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
如果李拿走前 3 颗，那么剩下的是 [4,5]，亚历克斯拿走后 5 颗赢得 10 分。
如果李拿走后 5 颗，那么剩下的是 [3,4]，亚历克斯拿走后 4 颗赢得 9 分。
这表明，取前 5 颗石子对亚历克斯来说是一个胜利的举动，所以我们返回 true 。

提示：
2 <= piles.length <= 500
piles.length 是偶数。
1 <= piles[i] <= 500
sum(piles) 是奇数。

*/
package main

import (
	"fmt"
)

func main() {
	piles := []int{3, 2, 10, 4}
	fmt.Println(stoneGame(piles))
}

/*
无情解法
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了96.67%的用户
*/
func stoneGame1(piles []int) bool {
	return true
}

/*
动态规划
执行用时 :8 ms, 在所有 Go 提交中击败了28.85%的用户
内存消耗 :6.5 MB, 在所有 Go 提交中击败了10.00%的用户

使用二维动态规划数组，dp[i][j]记录了在p[i]到p[j]的序列中先选的人能够取得的比对手多的最大石头数
最优子结构：dp[0][1]=max(piles[0]-dp[1][1],piles[1]-dp[0][0])
状态转换方程：dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
*/
func stoneGame(piles []int) bool {
	l := len(piles)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		a := make([]int, l)
		dp[i] = a
		dp[i][i] = piles[i]
	}
	for j := 1; j < l; j++ {
		for i := j - 1; i >= 0; i-- {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	//fmt.Println(dp)
	return dp[0][l-1] > 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
