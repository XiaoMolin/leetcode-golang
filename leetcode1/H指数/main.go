// H指数 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{1, 0}
	fmt.Println(hIndex(citations))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.3 MB, 在所有 Go 提交中击败了72.73%的用户

排序，找到引用次数大于所剩余论文的次数也即是
至多有h篇论文至少引用了h次
*/
func hIndex(citations []int) int {
	res := 0
	sort.Ints(citations)
	l := len(citations)
	for i := 0; i < l; i++ {
		if citations[i] >= l-i {
			res = l - i
			return res
		}
	}
	return 0
}
