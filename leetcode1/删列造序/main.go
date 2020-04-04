// 删列造序 project main.go
package main

import (
	"fmt"
)

func main() {
	A := []string{"cba", "daf", "ghi"} //1
	//A := []string{"a", "b"} //0
	//A := []string{"zyx", "wvu", "tsr"} //3
	//A := []string{"cba", "daf", "ghi"}
	//A := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(A))
}

/*
执行用时 :12 ms, 在所有 Go 提交中击败了67.65%的用户
内存消耗 :6.1 MB, 在所有 Go 提交中击败了86.36%的用户

遍历每一列只要有前一个大于后面的则删除的列数++
即用一个target标记最小的，用冒泡往后比较，只要target不移动则删除
时间复杂度 O(n*n)
空间复杂度 O(1)
*/
func minDeletionSize(A []string) int {
	res := 0
	x := len(A[0])
	y := len(A)
	for i := 0; i < x; i++ {
		target := A[0][i]
		for j := 0; j < y; j++ {
			if A[j][i] >= target {
				target = A[j][i]
			} else {
				res++
				break
			}
		}
	}
	return res
}
