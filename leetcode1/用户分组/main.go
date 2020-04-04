// 用户分组 project main.go
package main

import (
	"fmt"
)

func main() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	//groupSizes:= []int{2,1,3,3,3,2}
	//groupSizes:= []int{3,3,3,3,3,1,3}
	//groupSizes:= []int{3,3,3,3,3,1,3}
	fmt.Println(groupThePeople(groupSizes))
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了95.52%的用户
内存消耗 :6 MB, 在所有 Go 提交中击败了9.09%的用户

首先创建一个map，存放key=groupSizes[i],value=i
会得到一个像[1:[5] 3:[0 1 2 3 4 6]]这样的map
然后在创建一个数组res用来存放结果
遍历map，每一个value/key的长度 i++
res[i]中将res[i][j]=对应的value

*/
func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int, 0)
	for i, v := range groupSizes {
		arr, ok := m[v]
		if false == ok {
			arr = make([]int, 0)
		}
		arr = append(arr, i)
		m[v] = arr
	}
	fmt.Println(m)
	result := make([][]int, 0)
	for k, v := range m {
		for i := 0; i < len(v)/k; i++ {
			arr := make([]int, k)
			for j := 0; j < k; j++ {
				arr[j] = v[i*k+j]
			}
			result = append(result, arr)
		}
	}
	return result
}

/*
优化：
执行用时 :8 ms, 在所有 Go 提交中击败了95.52%的用户
内存消耗 :6 MB, 在所有 Go 提交中击败了23.64%的用户

*/
func groupThePeople1(groupSizes []int) [][]int {
	record := make(map[int][]int)
	var res [][]int
	for i := 0; i < len(groupSizes); i++ {
		record[groupSizes[i]] = append(record[groupSizes[i]], i)
	}
	for k, v := range record {
		var tmp []int
		for i := 0; i < len(v); i++ {
			tmp = append(tmp, v[i])
			if (i+1)%k == 0 {
				res = append(res, tmp)
				tmp = make([]int, 0)
			}
		}
	}
	return res
}
