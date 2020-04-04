// 煎饼排序 project main.go
package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 4, 1}
	fmt.Println(pancakeSort(a))
}

/*
3
4231
4
1324
2
3124
3
2134
2
1234
*/

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.4 MB, 在所有 Go 提交中击败了42.11%的用户
每次找到最大数的下标，若不为0则翻转从0到最大数下标的所有数
此时最大数在下标0，在翻转一次将最大数反转到最后一个
*/
func pancakeSort(a []int) []int {
	l, g := len(a), len(a)
	res := []int{}
	for i := 0; i < l; i++ {
		subscript := max(a)
		flip(a, subscript)
		res = append(res, subscript+1)
		flip(a, g-1)
		a = a[:g-1]
		res = append(res, g)
		g--
	}
	return res
}

func max(a []int) int {
	m := a[0]
	res := 0
	for i := 1; i < len(a); i++ {
		if m < a[i] {
			res = i
			m = a[i]
		}
	}
	return res
}

func flip(a []int, l int) {
	i := 0
	l = l
	for i < l {
		a[i], a[l] = a[l], a[i]
		i++
		l--
	}
}
