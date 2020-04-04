// 最大数 project main.go
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//nums := []int{10, 2} //210
	nums := []int{3, 30, 34, 5, 9} //9534330
	//nums:=[]int{}//
	fmt.Println(largestNumber(nums))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.7 MB, 在所有 Go 提交中击败了42.42%的用户
*/
/*
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	ns := Nums(nums)
	sort.Sort(ns)
	ret := ""
	for i := len(nums) - 1; i >= 0; i-- {
		ret = ret + strconv.Itoa(nums[i])
	}
	if ret[0] == '0' {
		return "0"
	}
	return ret
}

type Nums []int

func (n Nums) Len() int {
	return len(n)
}

func (n Nums) Less(a, b int) bool {
	na := n[a]
	nb := n[b]

	nab := strconv.Itoa(na) + strconv.Itoa(nb)
	nba := strconv.Itoa(nb) + strconv.Itoa(na)
	if nab > nba {
		return false
	}
	return true
}

func (n Nums) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
	return
}

*/

/*
执行用时 :4 ms, 在所有 Go 提交中击败了80.75%的用户
内存消耗 :2.8 MB, 在所有 Go 提交中击败了28.28%的用户
*/
func largestNumber(nums []int) string {
	res := []string{}
	for _, v := range nums {
		res = append(res, strconv.Itoa(v)) //int->string
	}
	sort.Strings(res)
	fmt.Println(res)
	for i := len(nums) - 1; i > 0; i-- {
		if res[i]+res[i-1] < res[i-1]+res[i] {
			res[i], res[i-1] = res[i-1], res[i]
			i += 2
			if i >= len(res) {
				i = len(res)
			}
		}
	}
	result := ""
	for i := len(res) - 1; i >= 0; i-- {
		result += res[i]
	}
	if result[:1] == "0" {
		return "0"
	}
	return result
}
