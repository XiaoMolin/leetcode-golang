// 加油站 project main.go
package main

import (
	"fmt"
)

func main() {
	//gas := []int{1, 2, 3, 4, 5}
	//cost := []int{3, 4, 5, 1, 2} //-2 -2 -2 3 3
	//gas := []int{2, 3, 4}
	//cost := []int{3, 4, 3} //-1 -1 1
	//gas := []int{5, 8, 2, 8}
	//cost := []int{6, 5, 6, 6} //-1 3 -4 2
	gas := []int{1, 2, 3, 4, 5, 5, 70}
	cost := []int{2, 3, 4, 3, 9, 6, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

/*
执行用时 :16 ms, 在所有 Go 提交中击败了38.46%的用户
内存消耗 :3.1 MB, 在所有 Go 提交中击败了16.35%的用户

将没一个站的得失油量计算，然后遍历只要从一个站开始
之后的每一个站油量都大于等于零则没问题
*/

func canCompleteCircuit1(gas []int, cost []int) int {
	l := len(gas)
	list := make([]int, l)
	res := 0
	for i := 0; i < l; i++ {
		list[i] = gas[i] - cost[i]
		res += list[i]
	}
	if res < 0 {
		return -1
	}
	for i := 0; i < l; i++ {
		res = 0
		if list[i] >= 0 {
			for j := i; j < l; j++ {
				res += list[j]
				if res < 0 {
					break
				}
				if j == l-1 {
					res = i
					return res
				}
			}
		}
	}
	return res
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了94.23%的用户
内存消耗 :3 MB, 在所有 Go 提交中击败了100.00%的用户


遍历一遍找到一个油量为最少的站，油量可能为负数，那么下一个站就是最有可能的站
如果遍历一遍总油量小于零，那么失败，若大于零，则为最少油的下一站


如果油是最少的，那么肯定下一个站油是增多的，
*/
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	min := gas[0] - cost[0] + 1
	start := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < min {
			min = sum
			fmt.Println(min)
			start = (i + 1) % n
		}
	}
	if sum < 0 {
		return -1
	}
	return start
}
