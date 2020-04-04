// 规划兼职工作 project main.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	fmt.Println(jobScheduling(startTime, endTime, profit))
}

/*
执行用时 :124 ms, 在所有 Go 提交中击败了56.67%的用户
内存消耗 :9.3 MB, 在所有 Go 提交中击败了27.27%的用户

dp动态规划

先将三个数组组合再一次，在以结束时间升序
type work struct {
	start int
	end   int
	price int
}
在使用一个数组记录做每一个工作之前最后能做的工作的下标
也就是说例如[{1 3 50} {2 4 10} {3 5 40} {3 6 70}]
第一个工作为-1，因为之前没有工作可做
第二个工作为-1，因为之前的工作结束时间在第二个工作之后即3>2
第三个工作为 0，因为只有第一个工作结束在第三个工作之前
第四个工作为 0，同理第三个工作
即prev = [-1,-1,0,0]

动态规划遍历一遍所有工作
假设做一个工作则，则判断prev是否大于等于0，
若大于等于零则，do=dp[prev[i]]+arr[i].price
表示做此工作前可以做第prev[i]个工作，
若小于零，则表示做此工作之前无法做任何工作do= arr[i].price
假设不做这个工作，则ndo=dp[i-1],表示收益为前状态的收益
然后比较一下做或是不做的收益大小，决定本次状态做还是不做
dp[i] = int(math.Max(float64(ndo), float64(do)))

最后的一个dp就是最终最大的收益
*/
type work struct {
	start int
	end   int
	price int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	l := len(startTime)
	arr := make([]work, 0, l)
	for i := 0; i < l; i++ {
		arr = append(arr, work{start: startTime[i], end: endTime[i], price: profit[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].end < arr[j].end
	})
	fmt.Println(arr)
	prev := make([]int, l)
	for i := 0; i < l; i++ {
		prev[i] = -1
	}
	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[i].start >= arr[j].end {
				prev[i] = j
				break
			}
		}
	}
	dp := make([]int, l)
	dp[0] = arr[0].price
	for i := 1; i < l; i++ {
		ndo := dp[i-1]
		do := arr[i].price
		if prev[i] >= 0 {
			do = dp[prev[i]] + arr[i].price
		}
		dp[i] = int(math.Max(float64(ndo), float64(do)))
	}
	return dp[l-1]
}
