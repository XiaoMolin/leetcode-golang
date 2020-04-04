// 最长回文子串 project main.go
package main

import (
	"fmt"
)

func main() {
	s := "asdasfasaaaaaasdasdasfawas"
	fmt.Println(longestPalindrome_2(s))
}

/*
暴力破解
执行用时 :1240 ms, 在所有 Go 提交中击败了5.04%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了91.67%的用户
*/
func isPalidromic(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome_0(s string) string {
	result := ""
	max := 0
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			test := s[i:j]
			if isPalidromic(test) && len(test) > max {
				result = test
				max = len(test)
			}
		}
	}
	return result
}

/*
中心扩展法
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了91.67%的用户
*/
func longestPalindrome_1(s string) string {
	if len(s) < 2 {
		return s
	}
	maxlen := 0
	begin := 0
	for i := 0; i < len(s); {
		if len(s)-i <= maxlen/2 {
			break
		}
		n, m := i, i
		for n < len(s)-1 && s[n] == s[n+1] {
			n++
		}
		i = n + 1
		for n < len(s)-1 && m > 0 && s[m-1] == s[n+1] {
			n++
			m--
		}
		temp := n + 1 - m
		if temp > maxlen {
			begin = m
			maxlen = temp
		}
	}
	return s[begin : begin+maxlen]
}

/*
最长公共子串-动态规划
执行用时 :136 ms, 在所有 Go 提交中击败了18.43%的用户
内存消耗 :17 MB, 在所有 Go 提交中击败了6.22%的用户
*/
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func longestPalindrome_2(s string) string {
	if s == "" {
		return s
	}
	s_ := Reverse(s)
	l := len(s)
	arr := make([][]int, l)
	for i := 0; i < l; i++ {
		arr[i] = make([]int, l)
	}
	maxlen := 0
	maxend := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if s[i] == s_[j] {
				if i == 0 || j == 0 {
					arr[i][j] = 1
				} else {
					arr[i][j] = arr[i-1][j-1] + 1
				}
			}
			if arr[i][j] > maxlen {
				beforeRev := l - 1 - j
				if beforeRev+arr[i][j]-1 == i {
					maxlen = arr[i][j]
					maxend = i
				}
			}
		}
	}
	res := s[maxend-maxlen+1 : maxend+1]
	return res
}
