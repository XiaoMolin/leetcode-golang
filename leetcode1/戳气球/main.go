// 戳气球 project main.go
package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}

/*


执行用时 :4 ms, 在所有 Go 提交中击败了91.57%的用户
内存消耗 :2.9 MB, 在所有 Go 提交中击败了58.70%的用户
以23125为例
因为只剩下一个气球时，爆掉后两边都要乘以1因此在数组两边添加数字1对于结果本身并没有影响
1 2 3 1 2 5 1
用dp[i][j]表示第i到第j个气球所能爆出的最大值
最优子结构：

...................................
 for(int r=2;r<n;r++)            //r为区间长度
            for(int i=0;i<n-r;i++){    //i为左区间
                int j=i+r;            //j为右区间
                for(int k=i+1;k<j;k++)
                    dp[i][j]=max(dp[i][j],dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j]);
            }


*/

func maxCoins(nums []int) int {
	l := len(nums)
	s := make([]int, l+2)
	s[0], s[l+1] = 1, 1
	for i := 1; i < l+1; i++ {
		s[i] = nums[i-1]
	}
	fmt.Println(s)
	dp := make([][]int, l+2)
	for i := 0; i < l+2; i++ {
		dp[i] = make([]int, l+2)
	}

	for i := l; i >= 1; i-- {
		for j := i; j <= l; j++ {
			if i == j {
				dp[i][j] = s[i]
				fmt.Println(s[i], i, j)
			}
			for k := i; k <= j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k-1]+dp[k+1][j]+s[k]*s[i-1]*s[j+1])
			}
		}
		fmt.Println(dp)
	}

	return dp[1][l]
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
贪心算法，失败
*/

func maxCoins2(nums []int) int {
	res := 0
	for i := len(nums); i > 2; i-- {
		k := findmin(nums)
		res += nums[k] * nums[k+1] * nums[k-1]
		nums = append(nums[:k], nums[1+k:]...)
	}
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			res += nums[0]*nums[1] + nums[1]
		} else {
			res += nums[0]*nums[1] + nums[0]
		}
	}
	if len(nums) == 1 {
		res = nums[0]
	}
	return res
}

func findmin(nums []int) int {
	l := len(nums)
	k := 1
	min := nums[1]
	for i := 1; i < l-1; i++ {
		if min > nums[i] {
			min = nums[i]
			k = i
		}
	}
	return k
}
