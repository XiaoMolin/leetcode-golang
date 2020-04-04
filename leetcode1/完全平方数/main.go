// 完全平方数 project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 12
	fmt.Println(numSquares(n))
}

/*
dp[1]=1
dp[2]=2
dp[3]=3
dp[4]=2
dp[5]=dp[4]+dp[1]=2+1=3
dp[6]=dp[3]+dp[3]=1+1=2
dp[7]=3
dp[8]=2
dp[9]=1
*/

/*
执行用时 :588 ms, 在所有 Go 提交中击败了7.16%的用户
内存消耗 :5.7 MB, 在所有 Go 提交中击败了83.95%的用户
对于每一个i，找每一对下标值之和等于i的数中找最小
*/
func numSquares1(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	a := math.Sqrt(float64(n))
	b := int(a)
	for i := 1; i <= b; i++ {
		dp[i*i] = 1
	}
	for i := 2; i <= n; i++ {
		if dp[i] != 1 {
			min := 1000000
			for j, k := 1, i-1; j <= k; {
				if dp[j]+dp[k] < min {
					min = dp[j] + dp[k]
				}
				j++
				k--
			}
			dp[i] = min
		}
	}
	return dp[n]
}

/*
执行用时 :192 ms, 在所有 Go 提交中击败了13.33%的用户
内存消耗 :5.8 MB, 在所有 Go 提交中击败了80.86%的用户
动态规划对于每一个i减去最大的完全平方数的dp值加一
不使用math.Min时间可以缩减为44ms
*/
func numSquares2(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-j*j]+1)))
		}
	}
	return dp[n]
}

/*
BFS
执行用时 :68 ms, 在所有 Go 提交中击败了28.89%的用户
内存消耗 :6.4 MB, 在所有 Go 提交中击败了26.54%的用户
*/

func numSquares(n int) int {
	queue := make([]int, 0)
	visited := map[int]bool{
		0: true,
	}
	queue = append(queue, 0)

	distance := 0
	for len(queue) > 0 {
		distance++
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			for j := 1; j*j+curr <= n; j++ {
				next := j*j + curr
				if next == n {
					return distance
				}
				if next < n && !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
			}
		}
		queue = queue[size:]
	}

	return distance
}
