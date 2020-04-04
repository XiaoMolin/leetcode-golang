// 最接近原点的 K 个点 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	K := 1
	points := [][]int{{1, 3}, {-2, 2}}
	fmt.Println(kClosest(points, K))
}

/*
执行用时 :136 ms, 在所有 Go 提交中击败了67.11%的用户
内存消耗 :8 MB, 在所有 Go 提交中击败了6.38%的用户

排序可得
*/
type point struct {
	points []int
	sqrt   int
}

func kClosest(points [][]int, K int) [][]int {
	l := len(points)
	if l == 0 {
		return nil
	}
	if l <= K {
		return points
	}
	p := make([]point, 0, l)
	res := [][]int{}
	for _, v := range points {
		p = append(p, point{v, v[0]*v[0] + v[1]*v[1]})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].sqrt < p[j].sqrt
	})
	for _, v := range p {
		if K > 0 {
			res = append(res, v.points)
			K--
		}
	}
	return res
}
