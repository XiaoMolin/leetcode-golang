// 不同路径 project main.go
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？









*/
package main

import (
	"fmt"
)

func main() {
	m := 7
	n := 3
	fmt.Println(uniquePaths(m, n))
}

/*
分析:每一个格子所能到达的路径都是与之相邻且在上方和左方格字路径的总和
最优子结构dp[1][1]=dp[1][0]+dp[0][1]
状态转换方程dp[i][j]=dp[i-1][j]+dp[i][j-1]
边界:靠左边和上边的所有格字只有1个路径

执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了74.54%的用户

优化 可以使用两个值来维护dp数组
*/

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if i == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//fmt.Println(dp)
	return dp[n-1][m-1]
}
