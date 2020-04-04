// 插入区间 project main.go
package main

import (
	"fmt"
)

func main() {
	//intervals := [][]int{{1, 2}, {3, 4}, {6, 7}, {8, 10}, {12, 16}}
	//newInterval := []int{4, 8}
	intervals := [][]int{{1, 5}}
	newInterval := []int{0, 3}
	fmt.Println(insert(intervals, newInterval))
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了92.20%的用户
内存消耗 :4.8 MB, 在所有 Go 提交中击败了83.05%的用户


*/
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for _, v := range intervals {
		// 新的在后面，没交集
		if v[1] < newInterval[0] {
			res = append(res, v)
		} else if newInterval[1] < v[0] {
			res = append(res, newInterval)
			newInterval = v
		} else {
			newInterval[0] = minInt(v[0], newInterval[0])
			newInterval[1] = maxInt(v[1], newInterval[1])
		}
	}
	res = append(res, newInterval)
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
