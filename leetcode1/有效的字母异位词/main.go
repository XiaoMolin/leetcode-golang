// 有效的字母异位词 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	//s := "anagram"
	//t := "nagaram" //true
	s := "rat"
	t := "car" //false
	fmt.Println(isAnagram(s, t))
}

/*
执行用时 :12 ms, 在所有 Go 提交中击败了30.96%的用户
内存消耗 :3.2 MB, 在所有 Go 提交中击败了23.32%的用户

string转byte[]然后切片排序for一下比较
*/
func isAnagram(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	s1 := []byte(s)
	t1 := []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	for i, elem := range s1 {
		if elem != t1[i] {
			return false
		}
	}
	return true
}
