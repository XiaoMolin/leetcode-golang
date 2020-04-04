// 柠檬水找零 project main.go
/*
在柠檬水摊上，每一杯柠檬水的售价为 5 美元。
顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零
，也就是说净交易是每位顾客向你支付 5 美元。
注意，一开始你手头没有任何零钱。
如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
来源：力扣（LeetCode）
*/
package main

import (
	"fmt"
)

func main() {
	//bills := []int{5, 5, 5, 10, 20} //t
	//bills := []int{5, 5, 10} //t
	//bills := []int{10, 10} //f
	bills := []int{5, 5, 10, 10, 20} //f
	//bills:=[]int{5,5,5,10,20}
	fmt.Println(lemonadeChange(bills))
}

/*
执行用时 :16 ms, 在所有 Go 提交中击败了99.21%的用户
内存消耗 :5.8 MB, 在所有 Go 提交中击败了92.16%的用户

只要拿到的前是10元，那么5元--
拿到20元，如果有10元，那么10元--，5元--，如果没10元，5元-=3
判断5元是否大于等于0若小于则为false
时间复杂度O(n)
空间复杂度O(1)
*/
func lemonadeChange(bills []int) bool {
	money5, money10 := 0, 0
	if bills[0] > 5 {
		return false
	}
	for _, money := range bills {
		switch money {
		case 5:
			money5++
		case 10:
			money5--
			money10++
		case 20:
			if money10 > 0 {
				money10 -= 1
				money5 -= 1
			} else {
				money5 -= 3
			}
		}
		if money5 < 0 {
			return false
		}
	}
	return true
}
