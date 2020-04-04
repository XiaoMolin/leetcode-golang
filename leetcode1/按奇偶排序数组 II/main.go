// 按奇偶排序数组 II project main.go
package main

import (
	"fmt"
)

func main() {
	a := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(a))
}

/*
执行用时 :796 ms, 在所有 Go 提交中击败了5.31%的用户
内存消耗 :6.3 MB, 在所有 Go 提交中击败了65.15%的用户
*/
func sortArrayByParityII1(a []int) []int {
	l := len(a)
	for i := 0; i < l; i++ {
		if i%2 == 0 && a[i]%2 == 1 {
			for j := i + 1; j < l; j++ {
				if j%2 == 1 && a[j]%2 == 0 {
					exchange(a, i, j)
				}
			}
		}
		if i%2 == 1 && a[i]%2 == 0 {
			for j := i + 1; j < l; j++ {
				if j%2 == 0 && a[j]%2 == 1 {
					exchange(a, i, j)
				}
			}
		}
	}
	return a
}

func exchange(a []int, x, y int) []int {
	a[x], a[y] = a[y], a[x]
	return a
}

/*
执行用时 :48 ms, 在所有 Go 提交中击败了23.01%的用户
内存消耗 :6.2 MB, 在所有 Go 提交中击败了89.39%的用户

遍历一遍将比匹配的下标存在两个数组里，然后交换

时间复杂度O(N)
空间复杂度O(N)
*/
func sortArrayByParityII2(a []int) []int {
	l := len(a)
	mismatch1 := []int{}
	mismatch2 := []int{}
	for i := 0; i < l; i++ {
		if a[i]%2 == 0 && i%2 == 1 {
			mismatch1 = append(mismatch1, i)
		}
		if a[i]%2 == 1 && i%2 == 0 {
			mismatch2 = append(mismatch2, i)
		}
	}
	for i := 0; i < len(mismatch1); i++ {
		exchange(a, mismatch1[i], mismatch2[i])
	}
	return a
}

/*
执行用时 :32 ms, 在所有 Go 提交中击败了38.05%的用户
内存消耗 :6.3 MB, 在所有 Go 提交中击败了42.42%的用户

双指针法

时间复杂度O(N)
空间复杂度O(1)
*/
func sortArrayByParityII3(a []int) []int {
	l := len(a)
	j := 1
	for i := 0; i < l; i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2
			}
			exchange(a, i, j)
		}
	}
	return a
}

/*
执行用时 :28 ms, 在所有 Go 提交中击败了77.88%的用户
内存消耗 :6.2 MB, 在所有 Go 提交中击败了92.42%的用户

创建一个数组，只要是偶数就放0+2n位置只要是奇数就放1+2n位置

时间复杂度O(N)
空间复杂度O(N)
*/
func sortArrayByParityII(A []int) []int {
	arr := make([]int, len(A))
	c := 0
	d := 1
	for _, v := range A {
		if v%2 == 0 {
			arr[c] = v
			c = c + 2
		} else {
			arr[d] = v
			d = d + 2
		}
	}
	return arr
}
