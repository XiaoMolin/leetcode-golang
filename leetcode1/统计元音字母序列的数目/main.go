// 统计元音字母序列的数目 project main.go
/*
给你一个整数 n，请你帮忙统计一下我们可以按下述规则形成多少个长度为 n 的字符串：

字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
每个元音 'a' 后面都只能跟着 'e'
每个元音 'e' 后面只能跟着 'a' 或者是 'i'
每个元音 'i' 后面 不能 再跟着另一个 'i'
每个元音 'o' 后面只能跟着 'i' 或者是 'u'
每个元音 'u' 后面只能跟着 'a'
由于答案可能会很大，所以请你返回 模 10^9 + 7 之后的结果。



示例 1：

输入：n = 1
输出：5
解释：所有可能的字符串分别是："a", "e", "i" , "o" 和 "u"。
示例 2：

输入：n = 2
输出：10
解释：所有可能的字符串分别是："ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" 和 "ua"。
示例 3：

输入：n = 5
输出：68

*/
package main

import (
	"fmt"
)

func main() {
	n := 4
	fmt.Println(countVowelPermutation(n))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户

a e
e a/i
i a/e/o/u
o i/u
u/a

发现规律
以a为结尾的有e,i,u
以e为结尾的有a,i
以i为结尾的有e,o
以o为结尾的有i
以u为结尾的有i,o

类比：青蛙跳台阶。
最终返回的计数总和 = 各个字母结尾的数量，相加。
*/
func countVowelPermutation(n int) int {
	mod := 1000000007
	a, e, i, o, u := 1, 1, 1, 1, 1
	res := 0
	for n > 1 {
		aa := (e + i + u) % mod
		ee := (a + i) % mod
		ii := (e + o) % mod
		oo := (i) % mod
		uu := (i + o) % mod
		a, e, i, o, u = aa, ee, ii, oo, uu
		n--
	}
	res = res + a + e + i + o + u
	return res % 1000000007
}
