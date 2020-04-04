// 翻转图像 project main.go
package main

import (
	"fmt"
)

func main() {
	a := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(a))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了87.10%的用户
内存消耗 :3 MB, 在所有 Go 提交中击败了20.00%的用户
*/
func flipAndInvertImage1(a [][]int) [][]int {
	l1 := len(a)
	l2 := len(a[0])
	for i := 0; i < l1; i++ {
		b := make([]int, 0, l2)
		for j := l2 - 1; j >= 0; j-- {
			b = append(b, (a[i][j]+1)%2)
		}
		for k := 0; k < l2; k++ {
			a[i][k] = b[k]
		}
	}
	return a
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了87.10%的用户
内存消耗 :2.9 MB, 在所有 Go 提交中击败了75.56%的用户

优化
*/
func flipAndInvertImage2(a [][]int) [][]int {
	l1 := len(a)
	l2 := len(a[0])
	for i := 0; i < l1; i++ {
		for k := 0; k < l2/2; k++ {
			a[i][k], a[i][l2-k-1] = a[i][l2-k-1], a[i][k]
		}
		for k := 0; k < (l2+1)/2; k++ {
			a[i][k], a[i][l2-k-1] = (a[i][k]+1)%2, (a[i][l2-k-1]+1)%2
		}
	}
	return a
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.9 MB, 在所有 Go 提交中击败了75.56%的用户

遍历一半，只要相对应的相等
如101，第0个和第2个相等那么就会变成相反的000
第一个和第一个相等变成010
时间复杂度为O(n*m)
*/
func flipAndInvertImage(a [][]int) [][]int {
	l1 := len(a)
	l2 := len(a[0])
	for i := 0; i < l1; i++ {
		for k := 0; k < (l2+1)/2; k++ {
			if a[i][k] == a[i][l2-k-1] {
				a[i][k] = 1 - a[i][k]
				a[i][l2-k-1] = a[i][k]
			}
		}
	}
	return a
}
