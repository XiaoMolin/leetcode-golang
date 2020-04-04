// 只有两个键的键盘 project main.go
package main

import (
	"fmt"
)

func main() {
	n := 1000
	fmt.Println(minSteps(n))
}

/*
执行用时 :20 ms, 在所有 Go 提交中击败了46.15%的用户
内存消耗 :2.3 MB, 在所有 Go 提交中击败了17.86%的用户

动态规划
*/
func minSteps(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		for j := i / 2; j >= 1; j-- {
			if j == 1 {
				dp[i] = i
				break
			}
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}

	}
	//fmt.Println(dp)
	return dp[n]
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户

找质数因子
*/

func minSteps1(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res += i
			n /= i
		}
	}
	return res
}
