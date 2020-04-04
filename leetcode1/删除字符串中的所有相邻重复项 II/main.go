// 删除字符串中的所有相邻重复项 II project main.go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	s := "pbbcggttciiippooaais"
	k := 2
	fmt.Println(removeDuplicates(s, k))
}

func removeDuplicates(s string, k int) string {
	result := list.New()
	result.PushBack(string(s[0]))
	time := 1
	number := 1
	for i, j := 1, len(s); i < j; {
		a := result.Back()
		for time < k {
			if a.Value == string(s[i]) && a.Prev() != nil {
				time++
				a = a.Prev()
			} else {
				result.PushBack(string(s[i])) //入栈
				time = 1
				number++
				break //退出循环
			}
		}
		var prev *list.Element
		for e := result.Back(); e != nil && time > 1; e = prev {
			prev = e.Prev()
			result.Remove(e)
			time--
			number--
		}
		i++
	}
	var res string = ""
	for i, j := result.Front(), 0; i != nil; i = i.Next() {
		res += i.Value.(string)
		j++
	}
	num := 1
	a := res[0]
	for i := 1; i < k; i++ {
		if a == res[i] {
			num++
		}
	}
	if num == k {
		res = res[k:]
	}
	return res
}
