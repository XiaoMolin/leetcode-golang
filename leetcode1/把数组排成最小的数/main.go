// 把数组排成最小的数 project main.go
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(minNumber(nums))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.7 MB, 在所有 Go 提交中击败了100.00%的用户

自定义go自带排序
将数组用字符的方式排序 '3'+'5'<'5'+'3'因此3排前面
排序后按顺序添加到字符串即可

优化思路，改进排序算法
*/
func minNumber1(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	ns := Nums(nums)
	sort.Sort(ns)
	ret := ""
	for i := len(nums) - 1; i >= 0; i-- {
		ret = ret + strconv.Itoa(nums[i])
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
	if nab < nba {
		return false
	}
	return true
}

func (n Nums) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
	return
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户

快排
*/
func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	qsort(nums)
	fmt.Println(nums)
	ret := ""
	for i := len(nums) - 1; i >= 0; i-- {
		ret = ret + strconv.Itoa(nums[i])
	}
	return ret
}
func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		nab := strconv.Itoa(data[i]) + strconv.Itoa(mid)
		nba := strconv.Itoa(mid) + strconv.Itoa(data[i])
		if nab < nba {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}

	qsort(data[:head])
	qsort(data[head+1:])
}
