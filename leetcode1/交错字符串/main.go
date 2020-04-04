// 交错字符串 project main.go
package main

import (
	"fmt"
)

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	len3 := len(s3)
	if len1+len2 != len3 {
		return false
	}
	dp := make([][]bool, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]bool, len2+1)
		if i == 0 {
			dp[0][0] = true
		} else {
			dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
		}
	}
	for j := 1; j <= len2; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			dp[i][j] = (dp[i-1][j] && s3[i+j-1] == s1[i-1]) || (dp[i][j-1] && s3[i+j-1] == s2[j-1])
		}
	}
	return dp[len1][len2]
}
