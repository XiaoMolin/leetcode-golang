// 一和零 project main.go
/*
在计算机界中，我们总是追求用有限的资源获取最大的收益。
现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。
你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。
注意:
给定 0 和 1 的数量都不会超过 100。
给定字符串数组的长度不会超过 600。

示例 1:
输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
输出: 4
解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。

示例 2:
输入: Array = {"10", "0", "1"}, m = 1, n = 1
输出: 2
解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。


*/
package main

import (
	"fmt"
)

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}

/*
我们用 dp(i, j) 表示使用 i 个 0 和 j 个 1，最多能拼出的字符串数目，
那么状态转移方程为dp[i][j]=max(dp[i][j],1+dp[i-zero][j-one])

执行用时 :28 ms, 在所有 Go 提交中击败了63.38%的用户
内存消耗 :3.4 MB, 在所有 Go 提交中击败了52.00%的用户
*/
func findMaxForm1(strs []string, m int, n int) int {
	l := len(strs)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for s := 0; s < l; s++ {
		zero := 0
		one := 0
		for z := 0; z < len(strs[s]); z++ {
			if strs[s][z] == '0' {
				zero++
			} else {
				one++
			}
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zero][j-one])
			}
		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	ma := make(map[int]int)
	min := 1000
	b := 0
	res := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			_, ok := ma[j]
			if !ok {
				if len(strs[j]) <= min {
					b = j
					min = len(strs[j])
				}
			}
		}
		ma[b] = 1
		for c := 0; c < len(strs[b]); c++ {
			if strs[b][c] == '0' {
				m--
			} else {
				n--
			}
			if m < 0 || n < 0 {
				return res
			}
		}
		res++
	}
	return res
}
