// 通过删除字母匹配到字典里最长单词 project main.go
package main

import (
	"fmt"
)

func main() {
	s := "aewfafwafjlwajflwajflwafj"
	d := []string{}
	d = append(d, "apple")
	d = append(d, "ewaf")
	d = append(d, "awefawfwaf")
	d = append(d, "awef")
	d = append(d, "awefe")
	d = append(d, "ewafeffewafewf")
	fmt.Println(findLongestWord(s, d))
}

/*
执行用时 :16 ms, 在所有 Go 提交中击败了88.04%的用户
内存消耗 :6.6 MB, 在所有 Go 提交中击败了89.19%的用户

拿d中每一个字符串与s相比较，第一个不匹配那么s下标往后移，
找到第一个字符相等的位置，先比较剩余长度若此时d的剩余长度大于s的剩余长度，则肯定不匹配
若小于则继续比较直到d到末尾，与之前匹配的d比较，长的为我们要的
若一样长，直接比较字典顺序，小的是我们要的

*/
func findLongestWord(s string, d []string) string {
	if s == "" || len(d) == 0 {
		return ""
	}
	var res string
	for _, v := range d {
		// 如果字典字符串比s还要长,那么肯定匹配不上
		if len(v) > len(s) {
			continue
		}
		i := 0
		j := 0
		for i < len(v) && j < len(s) {
			if v[i] == s[j] {
				// 当匹配到一个字符后,如果后面的长度比被匹配的还要长,那么肯定匹配不上
				if len(v[i:]) > len(s[j:]) {
					break
				}
				i++
			}
			j++
		}
		// 如果i 等于v的长度,说明 v已经完全匹配了,那么说明它肯定是s的子集
		if i == len(v) {
			// 比较之前最符合的子集长度与此次的相比
			if len(v) == len(res) {
				// 直接按照 大小比字典顺序
				if v < res {
					res = v
				}
			} else if len(v) > len(res) {
				res = v
			}
		}
	}

	return res
}
