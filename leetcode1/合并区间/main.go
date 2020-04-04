// 合并区间 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{0, 3}, {0, 1}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}

/*
执行用时 :28 ms, 在所有 Go 提交中击败了17.40%的用户
内存消耗 :5.8 MB, 在所有 Go 提交中击败了11.19%的用户

排序，遍历数组如果当前区间的右区间小于下一个区间的左区间，那么不用变化

如果当前区间的右区间等于下一个区间的左区间，那么将下一个区间
的左区间变成当前区间的左区间，即intervals[i+1][0] = intervals[i][0]

如果当前区间的右区间大于下一个区间的左区间，那么先将下一个区间的左区间变成
当前区间的左区间，再判断当前区间的右区间是否大于下一个区间的右区间
例如 [2,6][3,5]那么下一个区间的右区间也将变成当前区间的左区间即[2,6][2,6]

*/
func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i := 0
	fmt.Println(intervals)
	if i < len(intervals)-1 {
		for {
			if i >= len(intervals)-1 {
				break
			}
			if intervals[i][1] < intervals[i+1][0] {
				res = append(res, intervals[i])
			} else if intervals[i][1] == intervals[i+1][0] {
				intervals[i+1][0] = intervals[i][0]
			} else {
				intervals[i+1][0] = intervals[i][0]
				if intervals[i+1][1] < intervals[i][1] {
					intervals[i+1][1] = intervals[i][1]
				}
			}
			i++
		}
	}
	res = append(res, intervals[i])
	return res
}
