// 位1的个数 project main.go
package main

import (
	"fmt"
)

func main() {
	var num uint32 = 1231256236
	fmt.Println(hammingWeight(num))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户

*/
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num &= (num - 1)
		res++
	}
	return res
}
