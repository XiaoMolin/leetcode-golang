// 根据身高重建队列 project main.go
package main

import (
	"fmt"
)

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
	//fmt.Println(Sort(people))
}

/*
执行用时 :24 ms, 在所有 Go 提交中击败了70.73%的用户
内存消耗 :5.9 MB, 在所有 Go 提交中击败了94.25%的用户

先对输入数组排序，h升序，k降序 从头循环遍历 当前这个人就是剩下未安排的人中最矮的人，
他的k值就代表他在剩余空位的索引值 如果有多个人高度相同，要按照k值从大到小领取索引值 示例：

[ 0, 1, 2, 3, 4, 5 ] [ 4, 4 ] 4
[ 0, 1, 2, 3, 5 ]    [ 5, 2 ] 2
[ 0, 1, 3, 5 ]       [ 5, 0 ] 0
[ 1, 3, 5 ]          [ 6, 1 ] 3
[ 1, 5 ]             [ 7, 1 ] 5
[ 1 ]                [ 7, 0 ] 1
[ [ 5, 0 ], [ 7, 0 ], [ 5, 2 ], [ 6, 1 ], [ 4, 4 ], [ 7, 1 ] ]
*/
func reconstructQueue(people [][]int) [][]int {
	Sort(people)
	l := len(people)
	var list []int
	res := make([][]int, l)
	for i := 0; i < l; i++ {
		list = append(list, i)
		res[i] = make([]int, 2)
	}
	for i := 0; i < l; i++ {
		index := list[people[i][1]]
		res[index][0], res[index][1] = people[i][0], people[i][1]
		for j := 0; j < len(list); j++ {
			if list[j] == index {
				list = append(list[:j], list[j+1:]...)
				break
			}
		}
	}
	return res
}

func Sort(people [][]int) [][]int {
	l := len(people)
	var s [2]int
	for i := 0; i < l; i++ {
		minindex := i
		for j := i + 1; j < l; j++ {
			if people[j][0] < people[minindex][0] {
				minindex = j
			}
			if people[j][0] == people[minindex][0] && people[j][1] > people[minindex][1] {
				minindex = j
			}
		}
		s[0], s[1] = people[i][0], people[i][1]
		people[i][0], people[i][1] = people[minindex][0], people[minindex][1]
		people[minindex][0], people[minindex][1] = s[0], s[1]
	}
	return people
}
