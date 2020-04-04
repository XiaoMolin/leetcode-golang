// 玩筹码 project main.go
package main

import (
	"fmt"
)

func main() {
	//chips := []int{1, 2, 3}
	chips := []int{2, 2, 2, 3, 3}
	fmt.Println(minCostToMoveChips(chips))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.1 MB, 在所有 Go 提交中击败了66.67%的用户

奇数移动到奇数、偶数移动到偶数是无消耗的，意思就是算奇数和偶数的个数而已。
时间复杂度 O(n)
空间复杂度 O(1)
*/
func minCostToMoveChips(chips []int) int {
	l := len(chips)
	res1, res2 := 0, 0
	for i := 0; i < l; i++ {
		if chips[i]%2 == 1 {
			res1++
		} else {
			res2++
		}
	}
	if res1 < res2 {
		return res1
	}
	return res2
}
