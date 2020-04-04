// 翻转矩阵后的得分 project main.go
package main

import (
	"fmt"
)

func main() {
	A := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	//A := [][]int{{0, 0, 1, 1}}
	fmt.Println(matrixScore(A))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.5 MB, 在所有 Go 提交中击败了14.29%的用户

先把第一列全部变为1，然后遍历每一列
只要0的个数大于1，则翻转
最后就算每一行数的大小相加得出结果

时间复杂度O(n*n)
空间复杂度O(1)
*/
func matrixScore(A [][]int) int {
	res := 0
	l := len(A)
	for i := 0; i < l; i++ {
		if A[i][0] == 0 {
			for j, k := 0, len(A[i]); j < k; j++ {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}
	for i := 1; i < len(A[0]); i++ {
		a, b := 0, 0
		for j := 0; j < len(A); j++ {
			if A[j][i] == 0 {
				a++
			} else {
				b++
			}
		}
		if a > b {
			for j := 0; j < len(A); j++ {
				if A[j][i] == 0 {
					A[j][i] = 1
				} else {
					A[j][i] = 0
				}
			}
		}
	}
	for i := 0; i < len(A); i++ {
		base := 1
		for j := len(A[i]) - 1; j >= 0; j-- {
			if A[i][j] == 1 {
				res += base
				base *= 2
			} else {
				base *= 2
			}
		}
	}
	fmt.Println(A)
	return res
}
