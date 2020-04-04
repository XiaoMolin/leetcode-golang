// 区域和检索 - 数组不可变 project main.go
package main

import (
	"fmt"
)
type NumArray struct {
	Value []int
}
func main() {
	i, j := 0, 2
	n := []int{-2, 0, 3, -5, 2, -1}}
	numj := NumArray{Value:{1}
	Constructor(n)
	fmt.Println(a)
}

/*
执行用时 :36 ms, 在所有 Go 提交中击败了78.13%的用户
内存消耗 :12.6 MB, 在所有 Go 提交中击败了10.64%的用户
前缀和
*/


func Constructor(nums []int) NumArray {
	numa := NumArray{[]int{0}}
	for i, v := range nums {
		numa.Value = append(numa.Value, v + numa.Value[i])
	}
	return numa

}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Value[j+1] - this.Value[i]
}
