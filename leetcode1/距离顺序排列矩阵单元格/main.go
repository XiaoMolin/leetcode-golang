// 距离顺序排列矩阵单元格 project main.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	R := 1
	C := 2
	r0 := 0
	c0 := 0
	fmt.Println(allCellsDistOrder(R, C, r0, c0))
}

/*
执行用时 :28 ms, 在所有 Go 提交中击败了76.81%的用户
内存消耗 :7 MB, 在所有 Go 提交中击败了33.33%的用户
*/
type node struct {
	loc []int
	dis int
}

type nodeArr []node

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	nodeArr := make(nodeArr, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			nodeArr = append(nodeArr, node{[]int{i, j}, getManhattan(r0, i, j, c0)})
		}
	}
	sort.Sort(nodeArr)
	result := make([][]int, 0)
	for i := 0; i < nodeArr.Len(); i++ {
		result = append(result, nodeArr[i].loc)
	}
	return result
}

func getManhattan(r1, r2, c1, c2 int) int {
	return int(math.Abs(float64(r1-r2))) + int(math.Abs(float64(c1-c2)))
}

func (n nodeArr) Len() int {
	return len(n)
}

func (n nodeArr) Less(i, j int) bool {
	return n[i].dis < n[j].dis
}

func (n nodeArr) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
