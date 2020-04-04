// 回文子串 project main.go
package main

import (
	"fmt"
)

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}

/*
abb
aaa

res:=0

abbbb
dp[0][0]=1
dp[0][1]=dp[1][0]
dp[0][4]=dp[1][3]&v[0]==v[4]
动态规划

执行用时 :208 ms, 在所有 Go 提交中击败了5.88%的用户
内存消耗 :30.4 MB, 在所有 Go 提交中击败了6.06%的用户
*/

func countSubstrings1(s string) int {
	length := len(s)
	res := 0
	dp := make([][]bool, length, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length, length)
	}
	fmt.Println(dp)
	for j := 0; j < length; j++ {
		for i := j; i >= 0; i-- {
			if s[i] == s[j] && ((j-i < 2) || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了86.36%的用户
*/
func countSubstrings(s string) int {
	l := len(s)
	res := 0
	for i := 0; i < l; i++ {
		for start, end := i, i; start >= 0 && end < l; {
			if s[start] == s[end] {
				res++
				start--
				end++
			} else {
				break
			}
		}
		for start, end := i, i+1; start >= 0 && end < l; {
			if s[start] == s[end] {
				res++
				start--
				end++
			} else {
				break
			}
		}
	}
	return res
}
