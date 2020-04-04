// 除数博弈 project main.go
/*
爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
选出任一 x，满足 0 < x < N 且 N % x == 0 。
用 N - x 替换黑板上的数字 N 。
如果玩家无法执行这些操作，就会输掉游戏。
只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。

示例 1：
输入：2
输出：true
解释：爱丽丝选择 1，鲍勃无法进行操作。

示例 2：
输入：3
输出：false
解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。

*/
package main

import (
	"fmt"
)

func main() {
	n := 2
	fmt.Println(divisorGame(n))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户

巧解
最终结果应该是占到 2 的赢，占到 1 的输；
若当前为奇数，奇数的约数只能是奇数或者 1，因此下一个一定是偶数；
若当前为偶数， 偶数的约数可以是奇数可以是偶数也可以是 1，因此直接减 1，则下一个是奇数；
因此，奇则输，偶则赢。直接:

*/
func divisorGame1(n int) bool {
	for n >= 4 {
		n -= 2
	}
	if n == 3 {
		return false
	}
	return true

	/*
			if N & 1 > 0 {
		        return false
		    }
		    return true
	*/
}

/*
执行用时 :12 ms, 在所有 Go 提交中击败了23.89%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了41.67%的用户

动态规划
从头都尾判断n是否为爱丽丝赢，首先
1爱丽丝输
2爱丽丝赢
3爱丽丝输
4找到一个比4小的数i 满足4能整除i且4-i这个数的结果是爱丽丝输的那么这个数一定能赢
因为数字3先手爱丽斯输，但是数字4的话爱丽丝先手赢
总结：每次找 i%j == 0 && !dp[i-j]直到最后一个dp[n]就是最终的答案
*/

func divisorGame(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	dp := make([]bool, n+1)
	dp[1], dp[2] = false, true
	for i := 3; i < n; i++ {
		dp[i] = false
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
