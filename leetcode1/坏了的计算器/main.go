// 坏了的计算器 project main.go
package main

import (
	"fmt"
)

/*
在显示着数字的坏计算器上，我们可以执行以下两种操作：

双倍（Double）：将显示屏上的数字乘 2；
递减（Decrement）：将显示屏上的数字减 1 。
最初，计算器显示数字 X。

返回显示数字 Y 所需的最小操作数。

*/
func main() {
	//X := 2
	//Y := 3 //2
	//X := 5
	//Y := 8 //2
	//X := 3
	//Y := 10 //3
	//X := 1024
	//Y := 1 //1023
	X := 1
	Y := 1000000000 //39
	fmt.Println(brokenCalc(X, Y))
}

/*
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :1.9 MB, 在所有 Go 提交中击败了66.67%的用户

逆向思维，让y靠近x，转换为y除以2，无法整除就+1再除以2
时间复杂度： O(logY)。
空间复杂度： O(1)。
*/
func brokenCalc(x int, y int) int {
	res := 0
	for x < y {
		if y%2 == 0 {
			y /= 2
			res++
		} else {
			y = (y + 1) / 2
			res += 2
		}
	}
	return res + x - y
}
