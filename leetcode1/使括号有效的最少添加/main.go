// 使括号有效的最少添加 project main.go
package main

import (
	"fmt"
)

func main() {
	S := "())"
	//S := "((("
	//S := "()"
	//S := "()))(("
	//S:="())"
	fmt.Println(minAddToMakeValid(S))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2 MB, 在所有 Go 提交中击败了88.46%的用户

两个计数值rightLoss，rightLoss
如果是左括号则rightLoss++，表示缺多少右括号
如果是右括号则rightLoss--，如果rightLoss<0，
则表示左括号已经匹配完，多余了右括号，多余的右括号无法
通过后面的阔号匹配，因此需要leftLoss++
最后只需计算rightLoss+leftLoss即可
*/
func minAddToMakeValid1(S string) int {
	rightLoss := 0
	leftLoss := 0
	for _, c := range S {
		if c == '(' {
			rightLoss++
		} else {
			if rightLoss > 0 {
				rightLoss--
			} else {
				leftLoss++
			}
		}
	}
	return leftLoss + rightLoss
}
