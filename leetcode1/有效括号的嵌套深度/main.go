// 有效括号的嵌套深度 project main.go
package main

import (
	"fmt"
)

func main() {
	//seq := "(()())"
	seq := "()(())()"
	//seq:=""
	//seq:=""
	fmt.Println(maxDepthAfterSplit(seq))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.9 MB, 在所有 Go 提交中击败了85.71%的用户

因为是有效阔号，所以从第一个左括号开始
假设第一个左括号为0那么下一个左括号就是1后面同理
因为假设了第一个左括号为0，那么第一个右括号也应该为0
*/

func maxDepthAfterSplit(seq string) []int {
	l, r := 0, 0
	res := make([]int, len(seq))
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			res[i] = l
			l = (l + 1) % 2
		} else {
			res[i] = r
			r = (r + 1) % 2
		}
	}
	return res
}
