// 最后一块石头的重量 project main.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	stones := []int{1, 2, 3, 4, 5}
	fmt.Println(lastStoneWeight(stones))
}

/*
0ms,2mb

排序按从大到小，取最大两个数相减，然后放到里面排序
重复操作，得出最后的数

时间复杂度O(n*n)
空间复杂度O(1)
*/
func lastStoneWeight(stones []int) int {
	l := len(stones)
	if l == 0 {
		return 0
	}

	for i := l - 1; i > 0; i-- {
		sort.Ints(stones)
		stones[i-1] = int(math.Abs(float64(stones[i] - stones[i-1])))
	}
	return stones[0]
}
