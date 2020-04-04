// 分割平衡字符串 project main.go
package main

import (
	"fmt"
)

func main() {
	//s := "RLRRLLRLRL"
	//s := "RLLLLRRRLR"
	s := "LLLLRRRR"
	fmt.Println(balancedStringSplit(s))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了70.00%的用户

变量r_target，l_target标记'R'和'L'的次数
遍历一遍字符串，遇到'R',r_target++ 遇到'L'，l_target++
判断只要r_target==l_target res++
返回res即可

*/
func balancedStringSplit(s string) int {
	r_target := 0
	l_target := 0
	res := 0
	l := len(s)
	if l < 2 {
		return 1
	}
	for i := 0; i < l; i++ {
		if s[i] == 'R' {
			r_target++
		}
		if s[i] == 'L' {
			l_target++
		}
		if r_target == l_target {
			res++
		}

	}
	return res
}
