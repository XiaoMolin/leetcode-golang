// 目标和 project main.go
package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	s := 1
	fmt.Println(findTargetSumWays(nums, s))
}

/*
执行用时 :1100 ms, 在所有 Go 提交中击败了5.10%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了78.13%的用户
递归
*/
var count int = 0

func findTargetSumWays1(nums []int, S int) int {
	calculate(nums, 0, 0, S)
	return count
}

func calculate(nums []int, i, sum, S int) {
	if i == len(nums) {
		if S == sum {
			count++
		}
	} else {
		calculate(nums, i+1, sum+nums[i], S)
		calculate(nums, i+1, sum-nums[i], S)
	}
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了80.10%的用户
内存消耗 :2.5 MB, 在所有 Go 提交中击败了46.88%的用户

动态规划
因为是给数字加上正负号，因此每个数字有两种情况
dp[i][j]代表i个数能组合成j的个数

例如dp[3][1]代表前三个数能组成1的种类
假设1 1 1 1 1那么有-1+1+1，+1-1+1，+1+1-1三种情况，那么dp[3][1]=3
同理dp[3][0]=0.dp[3][-1]=3,dp[3][2]=0,dp[3][3]=1等等
那么dp[3][1]是怎么得来的呢?

因为我们知道nums[3]=1,所以dp[3][1]是由dp[3-1][1-1]和dp[3-1][1+1]得来的
dp[3][1]+=dp[2][0]
dp[3][1]+=dp[2][2]
dp[2][0]=2,dp[2][2]=1,因此dp[3][1]=3
又因为数组的下标不能为负数，且题目提示数组的和不会超过1000，
那么可以设置dp[i]的长度为2001,1000表示0

因此状态转移方程为：
dp[i][j+nums[i]] = dp[i][j+nums[i]] + dp[i-1][j]
dp[i][j-nums[i]] = dp[i][j-nums[i]] + dp[i-1][j]
如何确定j+-nums[i]不超过边界呢？
只要dp[i-1][j]>0,因为数组的和不会超过1000,
dp[i-1][j]>0表示的正是前i个数的组合

最后dp[len(nums)][S+1000]即可得出数组和为目标数 S 的所有添加符号的方法数

时间复杂度O(n*m)
空间复杂度O(n*m)

优化：我们发现dp[i][j]只与j有关，而i是递增的
因此我们可以使用dp[j]来代替dp[i][j]
时间复杂度O(n*m)
空间复杂度O(n)
*/
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == S || -nums[0] == S {
			return 1
		} else {
			return 0
		}
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < S {
		return 0
	}
	// 1000是和得正负分
	var dp [20][2001]int

	// 初始化dp[0]
	dp[0][nums[0]+1000] += 1
	dp[0][-nums[0]+1000] += 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < 2001; j++ {
			if dp[i-1][j] > 0 { //等于零代表目前没有和为j的组合
				dp[i][j+nums[i]] = dp[i][j+nums[i]] + dp[i-1][j]
				dp[i][j-nums[i]] = dp[i][j-nums[i]] + dp[i-1][j]
			}
			/*if j-nums[i] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
			if j+nums[i] <= 2000 {
				dp[i][j] += dp[i-1][j+nums[i]]
			}*/
		}
	}

	return dp[len(nums)-1][S+1000]
}

/*
原问题等同于： 找到nums一个正子集和一个负子集，使得总和等于target
我们假设P是正子集，N是负子集 例如： 假设nums = [1, 2, 3, 4, 5]，target = 3，
一个可能的解决方案是+1-2+3-4+5 = 3 这里正子集P = [1, 3, 5]和负子集N = [2, 4]

那么让我们看看如何将其转换为子集求和问题：
                  sum(P) - sum(N) = target
sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
                       2 * sum(P) = target + sum(nums)
因此，原来的问题已转化为一个求子集的和问题： 找到nums的一个子集 P，
使得sum(P) = (target + sum(nums)) / 2

上面的公式已经证明target + sum(nums)必须是偶数
*/
func findTargetSumWays2(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < S {
		return 0
	}
	// sum 和 S 必须同为奇数或者偶数
	if (sum+S)%2 == 1 {
		return 0
	}

	target := (sum + S) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}

/*
1.289 5g

1.183 芯片

1.280 通信
*/
