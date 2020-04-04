// 最小路径和 project main.go
/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了93.02%的用户
内存消耗 :4.4 MB, 在所有 Go 提交中击败了62.23%的用户

动态规划二维数组
dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
*/

func minPathSum1(grid [][]int) int {
	l := len(grid)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j == 0 {
					dp[0][0] = grid[0][0]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
				}
			}
		}
	}
	return dp[l-1][len(grid[0])-1]
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了93.02%的用户
内存消耗 :3.9 MB, 在所有 Go 提交中击败了100.00%的用户
动态规划原地
*/
func minPathSum(grid [][]int) int {
	for i, h := range grid {
		for j, _ := range h {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i][j-1] < grid[i-1][j] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
