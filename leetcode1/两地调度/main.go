// 两地调度 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	//costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	costs := [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}
	fmt.Println(twoCitySchedCost(costs))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.6 MB, 在所有 Go 提交中击败了29.41%的用户

贪心算法 先将a全加一遍，排序把另一半多加的减掉
*/
func twoCitySchedCost(costs [][]int) int {

	res, cons := 0, []int{}
	for _, v := range costs {
		res += v[0]
		cons = append(cons, v[1]-v[0])
	}
	sort.Ints(cons)
	for i := 0; i < len(cons)/2; i++ {
		res += cons[i]
	}

	return res
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.5 MB, 在所有 Go 提交中击败了76.47%的用户

贪心算法 先排序，优先a剩下为b
*/
func twoCitySchedCost01(costs [][]int) int {
	res := 0
	l := len(costs)
	insertingSort(costs)
	for i := 0; i < l/2; i++ {
		res += costs[i][0]
	}
	for i := l / 2; i < l; i++ {
		res += costs[i][1]
	}
	return res
}

func insertingSort(costs [][]int) [][]int {
	for i := 1; i < len(costs); i++ {
		for j := 0; j < i; j++ {
			//swap
			if (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1]) {
				costs[i], costs[j] = costs[j], costs[i]
			}
		}
	}
	return costs
}
