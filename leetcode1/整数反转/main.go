// 整数反转 project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 1534236469 //1534236469 //-2147483648 //-2147483412 //1534236469
	fmt.Println(reverse(x))
}

//第一次尝试
/*
执行用时 :
4 ms, 在所有 Go 提交中击败了47.08%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了14.54%的用户
*/
func reverse1(x int) int {
	if x == 0 {
		return x
	}
	y := strconv.Itoa(x)
	symbol := 1
	if x < 0 {
		symbol = -1
		y = y[1:]
	}
	l := len(y) - 1
	var result string
	for l >= 0 {
		result += string(y[l])
		l--
	}
	for {
		i := 0
		if result[i] == '0' {
			result = result[1:]
		} else {
			break
		}
	}
	res, _ := strconv.Atoi(result)
	if symbol == -1 {
		res = -res
	}
	if res > 2147483648 || res <= (-2147483648) {
		return 0
	}
	return res
}

//第二次尝试
/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了99.87%的用户
*/

func reverse(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		res = res*10 + pop
	}
	if res > 2147483648 || res <= (-2147483648) {
		return 0
	}
	return res
}
