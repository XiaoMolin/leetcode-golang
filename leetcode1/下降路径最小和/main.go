// 下降路径最小和 project main.go
/*
给定一个方形整数数组 A，我们想要得到通过 A 的下降路径的最小和。
下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。
在下一行选择的元素和当前行所选元素最多相隔一列。

示例：
输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：12
解释：
可能的下降路径有：
[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
和最小的下降路径是 [1,4,7]，所以答案是 12。

*/
package main

import (
	"fmt"
)

func main() {
	//a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	a := [][]int{{-80, -13, 22}, {83, 94, -5}, {73, -48, 61}}
	fmt.Println(minFallingPathSum(a))
}

/*
dp[i][j]表示从(i,j)元素开始的下降路径最小和
dp[i][j]可以下降到dp[i+1][j-1],dp[i+1][j],dp[i+1][j-1]
由于是从第一行开始，所以答案为min(dp[0][j])
执行用时 :8 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :5.6 MB, 在所有 Go 提交中击败了100.00%的用户
*/

func minFallingPathSum(A [][]int) int {
	l1, l2 := len(A), len(A[0])
	if l1 == 1 {
		min := A[0][0]
		for j := 0; j < l2; j++ {
			if min > A[0][j] {
				min = A[0][j]
			}
		}
		return min
	}
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		if i == l1-1 {
			for j := 0; j < l2; j++ {
				dp[i][j] = A[i][j]
			}
		}
	}
	for i := l1 - 2; i >= 0; i-- {
		for j := 1; j < l2-1; j++ {
			//fmt.Println(minthree(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1]))
			dp[i][j] = minthree(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1]) + A[i][j]
		}
		dp[i][0] = A[i][0] + mintwo(dp[i+1][0], dp[i+1][1])
		dp[i][l2-1] = A[i][l2-1] + mintwo(dp[i+1][l2-1], dp[i+1][l2-2])
	}
	min := dp[0][0]
	for j := 0; j < l2; j++ {
		if dp[0][j] < min {
			min = dp[0][j]
		}
	}
	fmt.Println(dp)
	return min
}

func mintwo(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func minthree(i, j, k int) int {
	min := mintwo(i, j)
	return mintwo(min, k)
}
