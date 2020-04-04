// Z 字形变换 project main.go
/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/
package main

import (
	"fmt"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了79.20%的用户
内存消耗 :6.5 MB, 在所有 Go 提交中击败了18.03%的用户

发像规律
第0行的相邻元素都是相差n=(numRows - 1) * 2
第1行的相邻元素都是相差分别为 n-2 和2
同理可得每一行相邻元素相差 n-2*k 和 k （k为第几行）
时间复杂度：O(n)
空间复杂度：O(n)
*/
func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 || l <= numRows {
		return s
	}
	res := ""
	n := (numRows - 1) * 2
	k := n
	for i := 0; i < numRows; i++ {
		for j := i; j < l; {
			res += string(s[j])
			if k == n {
				j += k
			} else {
				j += k
				if j < l && k != 0 {
					res += string(s[j])
				}
				j = j + n - k
			}
		}
		k -= 2
	}
	return res
}

/*
优化使用byte存储最后在转换成string
执行用时 :4 ms, 在所有 Go 提交中击败了94.53%的用户
内存消耗 :3.9 MB, 在所有 Go 提交中击败了98.36%的用户
*/
func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 || l <= numRows {
		return s
	}
	res := make([]byte, l)
	n := (numRows - 1) * 2
	k := n
	a := 0
	for i := 0; i < numRows; i++ {
		for j := i; j < l; {
			res[a] = s[j]
			a++
			if k == n {
				j += k
			} else {
				j += k
				if j < l && k != 0 {

					res[a] = s[j]
					a++
				}
				j = j + n - k
			}

		}
		k -= 2
	}
	return string(res)
}

/*

P   A   H   N
A P L S I I G
Y   I   R

PAHNAPLSIIGYYIIRR

PAHNAPLSIIGYIR


1 2 3 4 5 6 7 8 9

1   5   9
2 4 6 8
3   7

1     7        13
2   6 8     12 14
3 5   9  11    15
4     10       16


1       9
2     8 10
3   7   11
4 6     12
5       13
*/
