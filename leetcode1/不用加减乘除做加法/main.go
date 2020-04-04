// 不用加减乘除做加法 project main.go
package main

import (
	"fmt"
	_ "strconv"
)

func main() {
	a, b := 13, 9
	fmt.Println(add(a, b))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func add(a int, b int) int {
	for b != 0 {
		tmp := a ^ b
		b = (a & b) << 1
		a = tmp
		fmt.Println(tmp, b, a)
	}
	return a
}

/*

1101
1001

 0100

10010

*/

/*
思路一：先将a和b转换成二进制字符串，可以使用for循环 a>>1&1得出每个
位置的字符，然后翻转字符串。
此时使用两个指针从后往前对比两个字符串的字符，直到其中一个字符串下标为0
for直到其中一个字符串下标为0
判断如果int(str1[i])^int(str2[i])^carry==0//carry进位
那么在加一层判断是否是因为进位为0还是三个都为0
str1[i]=='0'||str2[i]=='0'
如果int(str1[i])^int(str2[i])^carry==1
那么也需判断是否因为进位
然后在判断此时进位是否为1
推出for后将剩余的子符串加回
然后倒置，最后将之转换成int

*/
