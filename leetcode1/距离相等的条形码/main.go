// 距离相等的条形码 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	barcodes := []int{1, 2, 2}
	fmt.Println(rearrangeBarcodes(barcodes))
}

/*
执行用时 :76 ms, 在所有 Go 提交中击败了80.00%的用户
内存消耗 :6.7 MB, 在所有 Go 提交中击败了50.00%的用户

将barcodes的值放入map里并统计每个值的个数
将map里的数据存到切n片里然后重定义排序
按照排序结果隔空一个位置插入到ret数组里
*/
func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(barcodes); i++ {
		m[barcodes[i]]++
	}
	n := make([]Node, 0, len(m))
	for k, v := range m {
		n = append(n, Node{k, v})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].count > n[j].count
	})
	ret := make([]int, len(barcodes))
	start := 0
	for _, v := range n {
		val, count := v.val, v.count
		for count > 0 {
			ret[start] = val
			start += 2
			count--
			if start >= len(barcodes) {
				start = 1
			}
		}
	}
	return ret
}

type Node struct {
	val   int
	count int
}
