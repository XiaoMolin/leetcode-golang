// 字符串相加 project main.go
package main

import (
	"fmt"
)

func main() {
	num1 := "5"
	num2 := "5"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	var x byte = 0
	var result = []byte{}
	for l1 >= 0 || l2 >= 0 || x != 0 {
		if l1 >= 0 {
			x += (num1[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			x += (num2[l2] - '0')
			l2--
		}
		result = append(result, (x%10)+'0')
		x = x / 10
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
