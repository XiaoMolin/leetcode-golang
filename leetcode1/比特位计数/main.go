// 比特位计数 project main.go
/*
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:
输入: 2
输出: [0,1,1]

示例 2:
输入: 5
输出: [0,1,1,2,1,2]

进阶:
给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）
来执行此操作。

*/
package main

import (
	"fmt"
)

func main() {
	num := 5
	fmt.Println(countBits(num))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了96.59%的用户
内存消耗 :4.2 MB, 在所有 Go 提交中击败了71.13%的用户

时间复杂度：O(n)O(n)。
空间复杂度：O(n)O(n)。

动态规划
0:0		2:10
1:01	3:11
发现每2的倍数最高位+1

0:0		4:10
1:01	5:101
2:10	6:110
3:11	7:111
因此dp[i+j]=dp[j]+1
每一轮i*=2
*/
func countBits1(num int) []int {
	dp := make([]int, num+1)
	j := 0
	i := 1
	for i < num+1 {
		for j < i && i+j < num+1 {
			dp[i+j] = dp[j] + 1
			j++
		}
		j = 0
		i *= 2
	}
	return dp
}

/*
 这个方法很赞 去掉最低位的1,然后用这个数的1的数量+1 就是结果了
i&(i-1)的作用是将i的二进制表示中的最低位为1的改为0

执行用时 :4 ms, 在所有 Go 提交中击败了96.59%的用户
内存消耗 :4.2 MB, 在所有 Go 提交中击败了99.30%的用户

时间复杂度：O(n)O(n)。
空间复杂度：O(n)O(n)。
*/
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}
