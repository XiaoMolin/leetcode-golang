// 字符串转换整数 (atoi) project main.go
package main

import (
	"fmt"
	_ "math"
)

func main() {
	str := "-2147483648"
	fmt.Println(myAtoi(str))
}

const (
	int32Max = 1<<31 - 1
	int32Min = -1 << 31
)

/*
执行用时 :4 ms, 在所有 Go 提交中击败了54.63%的用户
内存消耗 :2.3 MB, 在所有 Go 提交中击败了83.52%的用户
*/
func myAtoi(str string) int {
	n := len(str)
	var start, end int
	symbol := 1
	for start = 0; start < n; start++ { //找到数字的开头位置
		if str[start] >= '0' && str[start] <= '9' {
			break
		} else if str[start] == '+' {
			start++
			break
		} else if str[start] == '-' {
			symbol = -1
			start++
			break
		} else if str[start] != ' ' {
			return 0
		}
	}
	for end = start; end < n; end++ {
		if str[end] < '0' || str[end] > '9' {
			break
		}
	}
	res := 0
	for k := start; k < end; k++ {
		cur := int(str[k] - '0')

		if symbol > 0 {
			res = res*10 + cur
			if res > int32Max {
				return int32Max
			}
		} else {
			res = res*10 - cur
			if res < int32Min {
				return int32Min
			}
		}
	}
	return res
}
