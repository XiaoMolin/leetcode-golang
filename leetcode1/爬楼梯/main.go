// 爬楼梯 project main.go
/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

*/
package main

import (
	"fmt"
)

/*
动态规划
假设有十节楼梯，那么第十节一定是从第八节或是第九节走上来的
那么f(10)=f(9)+f(8)
最优子结构：f(10)=f(9)+f(8)
状态转移方程：f(n)=f(n-1)+f(n-2)
边界 f(1)=1,f(2)=2

执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了56.85%的用户
*/
func main() {
	n := 9
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	a := 1
	b := 2
	if n <= 2 {
		return n
	}
	for i := 3; i <= n; i++ {
		b, a = a+b, b
	}
	return b
}
