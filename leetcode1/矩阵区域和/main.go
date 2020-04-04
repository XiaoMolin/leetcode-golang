// 矩阵区域和 project main.go
package main

import (
	"fmt"
)

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	fmt.Println(matrixBlockSum(mat, k))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

/*
执行用时 :12 ms, 在所有 Go 提交中击败了75.00%的用户
内存消耗 :5.9 MB, 在所有 Go 提交中击败了100.00%的用户
利用二维前缀和求解
*/
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	mat1 := make([][]int, m)

	//计算前缀和
	//即mat1[i][j]为以(0,0)和(i,j)为顶点的矩形之和
	for i := 0; i < m; i++ {
		mat1[i] = make([]int, n)
		copy(mat1[i], mat[i])
	}
	for i := 1; i < m; i++ {
		r, rp := mat1[i], mat1[i-1]
		for j := 0; j < n; j++ {
			r[j] += rp[j]
		}
	}
	for i := 0; i < m; i++ {
		r := mat1[i]
		for j := 1; j < n; j++ {
			r[j] += r[j-1]
		}
	}

	answer := make([][]int, m)
	for i := 0; i < m; i++ {
		ar := make([]int, n)
		for j := 0; j < n; j++ {
			//确定所求矩形边界i1左边界,i2右边界,j1上边界,j2下边界
			i1, j1, i2, j2 := max(i-k, 0), max(j-k, 0), min(i+k, m-1), min(j+k, n-1)
			v := mat1[i2][j2] //v表示从(0,0)到(i2,j2)矩形和
			//减去左边矩形
			if i1 > 0 {
				v -= mat1[i1-1][j2]
			}
			//减去上边矩形
			if j1 > 0 {
				v -= mat1[i2][j1-1]
			}
			//加上重复减去的矩形
			if i1 > 0 && j1 > 0 {
				v += mat1[i1-1][j1-1]
			}
			ar[j] = v
		}
		answer[i] = ar
	}
	return answer
}
