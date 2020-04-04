// 移除盒子 project main.go
/*
不懂
给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。
每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
当你将所有盒子都去掉之后，求你能获得的最大积分和。

示例 1：
输入:
[1, 3, 2, 2, 2, 3, 4, 3, 1]
输出:
23

解释:
[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 分)
----> [1, 3, 3, 3, 1] (1*1=1 分)
----> [1, 1] (3*3=9 分)
----> [] (2*2=4 分)


提示：盒子的总数 n 不会超过 100。

*/
package main

import (
	"fmt"
)

func main() {
	boxes := []int{1, 3, 2, 2, 2, 3, 4, 3, 1}
	fmt.Println(removeBoxes(boxes))
}

/*
dp[i][j][k]代表i到
*/
func removeBoxes(boxes []int) int {
	l := len(boxes)
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, l)
		for j := 0; j < l; j++ {
			dp[i][j] = make([]int, l)
		}
	}
	for i := 0; i < l; i++ {
		for k := 0; k < l; k++ {
			dp[i][i][k] = (1 + k) * (1 + k)
		}
	}

	for t := 1; t < l; t++ {
		for j := t; j < l; j++ {
			i := j - t
			for k := 0; k <= i; k++ {
				res := (1+k)*(1+k) + dp[i+1][j][0]
				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[i] {
						res = max(res, dp[i+1][m-1][0]+dp[m][j][k+1])
					}
				}
				dp[i][j][k] = res
			}
		}
	}
	fmt.Println(dp)
	return dp[0][l-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
