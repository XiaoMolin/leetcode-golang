// 拼车 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	//trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	//capacity := 4
	//trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	//capacity := 5 //TTTTTTTTTTTT
	//trips := [][]int{{2, 1, 5}, {3, 5, 7}}
	//capacity := 3
	//trips := [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}
	//capacity := 11
	//trips := [][]int{{3, 2, 8}, {4, 4, 6}, {10, 8, 9}}
	//capacity := 11 //TTTTTTTTTT
	//trips := [][]int{{7, 5, 6}, {6, 7, 8}, {10, 1, 6}}
	//capacity := 16 //FFFFFFFFFFF
	//trips := [][]int{{4, 5, 6}, {6, 4, 7}, {4, 3, 5}, {2, 3, 5}}
	//capacity := 13 //TTTTTTTTTTTTTTTT
	trips := [][]int{{9, 3, 4}, {9, 1, 7}, {4, 2, 4}, {7, 4, 5}}
	capacity := 23 //TTTTTTTTTTTTTTTT
	fmt.Println(carPooling(trips, capacity))
}

/*
执行用时 :676 ms, 在所有 Go 提交中击败了7.69%的用户
内存消耗 :5.9 MB, 在所有 Go 提交中击败了11.11%的用户
*/
func carPooling1(trips [][]int, capacity int) bool {
	// 按照上车的站台顺序排序
	for i := 0; i < len(trips); i++ {
		sort.Slice(trips, func(i, j int) bool {
			return trips[i][1] <= trips[j][1]
		})
	}
	car := make(map[int]int, capacity)
	total := 0 // 车上所有的人
	// 单次行程安排
	for i := 0; i < len(trips); i++ {
		if trips[i][0] > capacity {
			return false
		}
		// 在某一站stations乘客people先下车
		for stations, people := range car {
			if stations <= trips[i][1] {
				total -= people
				delete(car, stations)
			}
		}
		total += trips[i][0] // 在某一站stations乘客people上车之后的总人数
		if total > capacity {
			return false
		}
		car[trips[i][2]] += trips[i][0] // 需要在某一站下车的人数
	}
	return true
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了96.15%的用户
内存消耗 :3.7 MB, 在所有 Go 提交中击败了22.22%的用户
使用map记录每一个站上车和下车的人数差
若大于零则上车人数大于下车人数
遍历一次车站数座位数减去每一个站的上下车人数差，大于等于0则为true
小于零则为false
*/
func carPooling(trips [][]int, capacity int) bool {
	m := make(map[int]int)
	for _, trip := range trips {
		m[trip[1]] += trip[0] // 记录每一站台上车的乘客
		m[trip[2]] -= trip[0] // 记录每一站台下车的乘客
	}
	for i := 0; i <= 1000; i++ {
		capacity -= m[i] // 每一站台车子的空余座位数
		if capacity < 0 {
			return false
		}
	}
	return true
}
