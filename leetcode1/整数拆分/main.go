// 整数拆分 project main.go
/*
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。

示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。

说明: 你可以假设 n 不小于 2 且不大于 58。
*/
package main

import (
	"fmt"
)

func main() {
	n := 10
	fmt.Println(integerBreak(n))
}

/*
总结规律 贪心算法
假设n=i+j，若要i*j最大那么i与j越接近最好
那么如何让n=i+j+k+..使得他们的乘积最大呢
我们可以先将n=i+j拆分成i=k+z使得组成i的两个数使得i最大
又因为当一个数大于3时他总是可以拆分成两个乘积比自己大的数
且拆分成3比拆分成2更有优势
例如6=2+2+2 6=3+3但是2*2*2=8，3*3=9

执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func integerBreak1(n int) int {
	/*res := 1
	three := 0
	two := 0
	if n<=3{
		return n-1
	}
	if n%3 == 0 {
		three = n / 3
	}
	if n%3 == 1 {
		three = n/3 - 1
		two = 2
	}
	if n%3 == 2 {
		three = n / 3
		two = 1
	}
	for three > 0 {
		res *= 3
		three--
	}
	for two > 0 {
		res *= 2
		two--
	}
	return res*/
	//化简
	res := 1
	if n <= 3 {
		return n - 1
	}
	for n > 4 {
		n -= 3
		res *= 3
	}
	return res * n
}

/*
动态规划
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了32.00%的用户
显然dp[2]等于 1，外层循环从 3 开始遍历，一直到 n 停止。
内层循环 j 从 1 开始遍历，一直到 i 之前停止，它代表着
数字 i 可以拆分成 j + (i - j)。但 j * (i - j)不一定是最大乘积，
因为i-j不一定大于dp[i - j]（数字i-j拆分成整数之和的最大乘积），
这里要选择最大的值作为 dp[i] 的结果。

空间复杂度是 O(N)，时间复杂度是 O(N^2)
*/

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 0, 1, 2
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func Max(i, j, k int) int {
	if i > j && i > k {
		return i
	} else if j > k {
		return j
	} else {
		return k
	}
}
