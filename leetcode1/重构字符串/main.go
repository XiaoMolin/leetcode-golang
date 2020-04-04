// 重构字符串 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "aab"
	fmt.Println(reorganizeString(s))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了60.00%的用户

统计每个字符数量，排序降序，第一个次数大于长度一半则失败返回""
小于则按隔空插入
*/

type letter struct {
	c     byte
	count int
}

func reorganizeString(S string) string {
	N := len(S)
	s := []byte(S)
	m := make(map[byte]int)
	for _, v := range s {
		m[v]++
	}

	letters := make([]letter, 0, N)
	for k, v := range m {
		letters = append(letters, letter{k, v})
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i].count > letters[j].count
	})

	if letters[0].count > (N+1)/2 {
		return ""
	}
	res := make([]byte, N, N)
	index := 0
	for _, v := range letters {
		num := v.count
		for num > 0 {
			if index >= N {
				index = 1
			}
			res[index] = v.c
			index += 2
			num--
		}
	}

	return string(res)
}
