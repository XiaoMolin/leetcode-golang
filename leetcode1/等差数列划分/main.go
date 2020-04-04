// 等差数列划分 project main.go
/*
如果一个数列至少有三个元素，并且任意两个相邻元素之差相同，则称该数列为等差数列。

例如，以下数列为等差数列:
1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9
以下数列不是等差数列。
1, 1, 2, 5, 7

数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。
如果满足以下条件，则称子数组(P, Q)为等差数组：
元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。
函数要返回数组 A 中所有为等差数组的子数组个数。

示例:
A = [1, 2, 3, 4]
返回: 3, A 中有三个子等差数组: [1, 2, 3], [2, 3, 4] 以及自身 [1, 2, 3, 4]。

*/
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(numberOfArithmeticSlices(a))
}

/*
思路：
123为1
1234为1+2
12345为1+2+3
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了85.71%的用户
时间复杂度O(n)
空间复杂度O(1)
*/
func numberOfArithmeticSlices(a []int) int {
	l := len(a)
	res := 0
	for i := 0; i < l-2; i++ {
		j := i + 1
		k := i + 2
		value := 1
		difference := a[j] - a[i]
		for ; k < l; k++ {
			if a[k]-a[j] == difference {
				res += value
				value++
				j++
				i = k
			} else {
				break
			}
		}
	}
	return res
}
