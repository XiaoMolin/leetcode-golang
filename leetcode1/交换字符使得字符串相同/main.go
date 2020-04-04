// 交换字符使得字符串相同 project main.go
package main

import (
	"fmt"
)

func main() {
	s1 := "xx"
	s2 := "yy"
	//s1 := "xy"
	//s2 := "yx"
	//s1 := "xx"
	//s2 := "xy"
	//s1 := "xxyyxyxyxx"
	//s2 := "xyyxyxxxyx"
	fmt.Println(minimumSwap(s1, s2))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户

因为两个字符串长度相同，因此只要x或y有一个的总数为奇数，那么就无法使两个字符串相同
使用Parget计算x的个数，若为奇数返回-1
只要不为奇数那么就肯定可以使之相同，如果出现有双数对x和y不匹配时那么只要交换一次即可，如示例一
如果是单数对移动次数为n/2+n%2，因为除了x和y不匹配还有y和x不匹配，所以要计算两次
从头到尾只遍历一遍时间复杂度为O(n)空间复杂度为O(1)
*/

func minimumSwap(s1 string, s2 string) int {
	res := 0
	l := len(s1)
	Parity := 0
	target_x, target_y := 0, 0
	for i := 0; i < l; i++ {
		if s1[i] == 'x' {
			Parity++
		}
		if s2[i] == 'x' {
			Parity++
		}
		if s1[i] != s2[i] && s1[i] == 'x' {
			target_x++
		}
		if s1[i] != s2[i] && s1[i] == 'y' {
			target_y++
		}
	}
	if Parity%2 == 1 {
		return -1
	}
	res = target_x/2 + target_x%2 + target_y/2 + target_y%2
	return res
}
