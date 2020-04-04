// 划分字母区间 project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "ababcbacadefegdehijhklij"
	//S:=""
	//S:=""
	//S:=""
	fmt.Println(partitionLabels(S))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了50.00%的用户
内存消耗 :2.3 MB, 在所有 Go 提交中击败了22.73%的用户
*/
func partitionLabels1(S string) []int {
	res := make([]int, 0)
	s := make(map[byte]int)
	begin := 0
	substr := ""
	for i := range S {
		if _, ok := s[S[i]]; !ok {
			substr += string(S[i])
			s[S[i]] = 0
		}
		fmt.Println(substr, i)
		if !strings.ContainsAny(S[i+1:], substr) {
			res = append(res, i-begin+1)
			begin = i + 1
			substr = ""
			s = make(map[byte]int)
		}
	}
	return res
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了50.00%的用户
内存消耗 :2.4 MB, 在所有 Go 提交中击败了22.73%的用户

第一个for：
用map存放字母最后出现的下标
用list和substr分别存放字母最开始出现的下标和字母
第二个for：
因为每个字母在在map，substr和lit只出现一次且顺序相同，所以用list长度即可
遍历过程中找substr中字母所对应map里的位置，即最后出现位置
只要substr字母出现在之前即对比它最后下标是否大于之前的，大于则交换last = 较大的
只要出现字母出现的下标大于目前所锁定的字母最后出现下标大的情况，则分割
res=append(res,last+1-start)
*/
func partitionLabels(S string) []int {
	res := make([]int, 0)
	substr := ""
	s := make(map[string]int)
	var list []int
	for i := range S {
		if _, ok := s[string(S[i])]; !ok {
			s[string(S[i])] = i
			substr += string(S[i])
			list = append(list, i)
		} else {
			s[string(S[i])] = i
		}
	}
	//fmt.Println(s)
	//fmt.Println(list)
	//fmt.Println(substr)
	last := 0
	start := 0
	for i := 0; i < len(list); {
		if s[string(substr[i])] > last {
			last = s[string(substr[i])]
		}
		var j int
		if i < len(list)-1 {
			j = i + 1
		}
		if list[j] < last {
			i++
		} else {
			res = append(res, last+1-start)
			start = last + 1
			i++
		}
		//fmt.Println(i)
	}
	res = append(res, last-start+1)
	return res
}
