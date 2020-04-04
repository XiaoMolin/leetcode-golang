// 最长回文子序列 project main.go
package main

import (
	"fmt"
)

func main() {
	s := "bbabaasfasgasdawsdaasfasdawsaf"
	fmt.Println(longestPalindromeSubseq(s))
}

/*
动态规划
思路用bp[i][j]表示i到j之间的最大回文子序列
bp[i][i]=1
假设第3和第6个字符相等那么
bp[3][6]=bp[4][5]+2//也就是在3和6之间的两个下标之间的回文子串数加二
假设不等那么
bp[3][6]=max(bp[3][5],bp[2][6])
然后bp[3][5],bp[2][6]又是按照bp[3][6]的求法获得

状态转移方程：
bp[i][j] = bp[i+1][j-1] + 2
bp[i][j] = max(bp[i+1][j], bp[i][j-1])
发现每个i都需要后面的一个i的值
j则需要前一个j的值
那么我们的顺序是逆序遍历i
顺序遍历j，且j从i+1开始

如果顺序从0开始
那么bp[0][1]=bp[1][0]+2
那么到bp[0][3]的时候如果相等但是bp[1][2]我们还没有计算出来，那么肯定得不到正确的答案

执行用时 :28 ms, 在所有 Go 提交中击败了80.72%的用户
内存消耗 :17.3 MB, 在所有 Go 提交中击败了17.39%的用户
*/

func longestPalindromeSubseq1(s string) int {
	l := len(s)
	str := []byte(s)
	//fmt.Println(str, l)
	bp := make([][]int, l)
	for i := 0; i < l; i++ {
		a := make([]int, l)
		bp[i] = a
	}

	for i := l - 1; i >= 0; i-- {
		bp[i][i] = 1
		for j := i + 1; j < l; j++ {
			if str[i] == str[j] {
				bp[i][j] = bp[i+1][j-1] + 2
			} else {
				bp[i][j] = max(bp[i+1][j], bp[i][j-1])
			}
		}
	}

	fmt.Println(bp)
	return bp[0][l-1]

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestPalindromeSubseq(s string) int {
	N := len(s)
	dp := make([]int, N)
	f := make([]int, 2)
	for i, _ := range dp {
		dp[i] = 1
	}
	for i := N; i >= 0; i-- {
		f[0] = 0
		for j := i + 1; j < N; j++ {
			f[1] = dp[j]
			if s[i] == s[j] {
				dp[j] = f[0] + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			f[0] = f[1]
		}
	}
	return dp[N-1]
}
