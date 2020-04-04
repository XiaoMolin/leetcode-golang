// 三角形最小路径和 project main.go
/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/
package main

import (
	"fmt"
)

func main() {
	//triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	triangle := [][]int{{-1}, {2, 3}, {1, -1, -3}}
	fmt.Println(minimumTotal(triangle))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了96.42%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了88.17%的用户

动态规划
triangle[i][j] = triangle[i][j] + int(math.Min(float64(triangle[i-1][j]), float64(triangle[i-1][j-1])))
*/
func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			n := len(triangle[i]) - 1
			if j == 0 {
				triangle[i][0] = triangle[i][0] + triangle[i-1][0]
			} else if j == n {
				triangle[i][n] = triangle[i][n] + triangle[i-1][j-1]
			} else {
				triangle[i][j] = triangle[i][j] + int(math.Min(float64(triangle[i-1][j]), float64(triangle[i-1][j-1])))
			}
		}
	}

	res := math.MaxInt32
	n := len(triangle[len(triangle)-1]) - 1
	for i := 0; i <= n; i++ {
		if triangle[n][i] < res {
			res = triangle[n][i]
		}
	}
	return res
}
