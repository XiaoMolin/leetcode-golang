// 罗马数字转整数 project main.go
package main

import (
	"fmt"
)

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}

/*
执行用时 :44 ms, 在所有 Go 提交中击败了5.58%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了97.68%的用户
*/
func romanToInt1(s string) int {
	l := len(s) - 1
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	a, b, c := 0, 0, 0
	for l >= 0 {
		if _, ok := m[string(s[l])]; ok {
			if a == 1 && m[string(s[l])] == 1 {
				res -= 2
			}
			if b == 1 && m[string(s[l])] == 10 {
				res -= 20
			}
			if c == 1 && m[string(s[l])] == 100 {
				res -= 200
			}
			a = 0
			b = 0
			c = 0
			res += m[string(s[l])]
			if m[string(s[l])] == 5 || m[string(s[l])] == 10 {
				a = 1
			}
			if m[string(s[l])] == 50 || m[string(s[l])] == 100 {
				b = 1
			}
			if m[string(s[l])] == 500 || m[string(s[l])] == 1000 {
				c = 1
			}
			l--
		}
	}
	return res
}

/*
执行用时 :8 ms, 在所有 Go 提交中击败了81.20%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了97.68%的用户

如果前面小于后面那么只有4，9，40，90，400，900这几种情况
每次判断两个小于就先加在减，小于就加
*/
func romanToInt(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	a := [...]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	rst := 0
	for i := 0; i < sLen; i++ {
		if i+1 < sLen && a[s[i+1]] > a[s[i]] {
			rst += a[s[i+1]] - a[s[i]]
			i++
		} else {
			rst += a[s[i]]
		}
	}

	return rst
}
