// 石子游戏 II project main.go
/*
亚历克斯和李继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
亚历克斯和李轮流进行，亚历克斯先开始。最初，M = 1。
在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
游戏一直持续到所有石子都被拿走。
假设亚历克斯和李都发挥出最佳水平，返回亚历克斯可以得到的最大数量的石头。
示例：

输入：piles = [2,7,9,4,4]
输出：10
解释：
如果亚历克斯在开始时拿走一堆石子，李拿走两堆，接着亚历克斯也拿走两堆。
在这种情况下，亚历克斯可以拿到 2 + 4 + 4 = 10 颗石子。
如果亚历克斯在开始时拿走两堆石子，那么李就可以拿走剩下全部三堆石子。
在这种情况下，亚历克斯可以拿到 2 + 7 = 9 颗石子。
所以我们返回更大的 10。

提示：
1 <= piles.length <= 100
1 <= piles[i] <= 10 ^ 4


*/
package main

import (
	"fmt"
)

func main() {
	//piles := []int{2, 7, 9, 4, 4}
	piles := []int{1, 2, 3, 4, 5, 100}
	fmt.Println(stoneGameII(piles))
}

/*
动态规划
用dp[i][m]记录从当前石头开始到最后一个石头，M为m的情况下能获得最多的石头数
通过自底向上的思路首先计算dp[4][1]=4
最优子结构：
dp[3][1]=max(dp[3][1],sum[5]-sum[3]-dp[3+j][max(m,j)])
dp[3][2]=max(dp[3][2],sum[5]-sum[3]-dp[3+j][max(m,j)])
。。。
dp[2][1]=max(dp[2][1],sum[5]-sum[2]-dp[2+j][max(m,j)])
dp[2][2]=max(dp[2][2],sum[5]-sum[2]-dp[2+j][max(m,j)])
状态转换方程：
dp[i][m]=max(dp[i][m],sum[l]-sum[i]-dp[i+j][max(m,j)])

执行用时 :4 ms, 在所有 Go 提交中击败了61.54%的用户
内存消耗 :3.2 MB, 在所有 Go 提交中击败了33.33%的用户

时间复杂度O(n*n)
空间复杂度O(n*n)
*/
func stoneGameII(piles []int) int {
	l := len(piles)
	dp := make([][]int, l+1)
	for i := 0; i < l+1; i++ {
		a := make([]int, l+1)
		dp[i] = a
	}
	sum := make([]int, l+1)
	for i := 1; i < l+1; i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}
	fmt.Println(sum)
	dp[l-1][1] = piles[l-1]
	for i := l; i >= 0; i-- {
		for m := 1; m <= i || m == 1; m++ {
			for j := 1; j <= (2 * m); j++ {
				if i+j > len(piles) {
					break
				} else {
					dp[i][m] = max(dp[i][m], sum[l]-sum[i]-(dp[i+j][max(m, j)]))
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[0][1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func stoneGameII2(piles []int) int {
	n := len(piles)
	var dp [101][101]int
	sum := 0
	for i := n - 1; i >= 0; i-- {
		sum += piles[i]
		for j := n; j >= 0; j-- {
			if i+2*j >= n {
				dp[i][j] = sum
			} else {
				for x := 1; i+x < n && x <= 2*j; x++ {
					dp[i][j] = max(dp[i][j], sum-dp[i+x][max(j, x)])
				}
			}
		}
	}
	return dp[0][1]
}
