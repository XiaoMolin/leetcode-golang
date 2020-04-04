// 将矩阵按对角线排序 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	fmt.Println(diagonalSort(mat))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了92.86%的用户
内存消耗 :5 MB, 在所有 Go 提交中击败了100.00%的用户

从右上开始按对角线遍历存放到ar数组里，然后排序再放回mat
到中间时，使用另一个for从第二列开始直到最后一列
*/

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	l := min(m, n)
	for i := n - 1; i >= 0; i-- {
		ar := make([]int, 0, l)
		r, c := 0, i
		for r < m && c < n {
			ar = append(ar, mat[r][c])
			r, c = r+1, c+1
		}
		sort.Ints(ar)
		r, c, k := 0, i, 0
		for r < m && c < n {
			mat[r][c] = ar[k]
			r, c, k = r+1, c+1, k+1
		}
	}
	for i := 1; i < m; i++ {
		ar := make([]int, 0, l)
		r, c := i, 0
		for r < m && c < n {
			ar = append(ar, mat[r][c])
			r, c = r+1, c+1
		}
		sort.Ints(ar)
		r, c, k := i, 0, 0
		for r < m && c < n {
			mat[r][c] = ar[k]
			r, c, k = r+1, c+1, k+1
		}
	}
	return mat
}
