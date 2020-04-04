// 柱状图中最大的矩形 project main.go
package main

import (
	"fmt"
)

func main() {
	heights := []int{4, 2}
	fmt.Println(largestRectangleArea(heights))
}

/*
执行用时 :496 ms, 在所有 Go 提交中击败了32.45%的用户
内存消耗 :4.4 MB, 在所有 Go 提交中击败了91.78%的用户

暴力破解
时间复杂度O(N*N)
*/
func largestRectangleArea1(heights []int) int {
	res := 0
	l := len(heights)
	for i := 0; i < l; i++ {
		min := heights[i]
		j, k := i, i
		for j >= 0 && heights[j] >= min {
			j--
		}
		for k < l && heights[k] >= min {
			k++
		}
		a := heights[i] * (k - j - 1)
		if a > res {
			res = a
		}
	}
	return res
}

/*
执行用时 :724 ms, 在所有 Go 提交中击败了17.70%的用户
内存消耗 :9.1 MB, 在所有 Go 提交中击败了7.53%的用户

分治
通过观察，可以发现，最大面积矩形存在于以下几种情况：
确定了最矮柱子以后，矩形的宽尽可能往两边延伸。
在最矮柱子左边的最大面积矩形（子问题）。
在最矮柱子右边的最大面积矩形（子问题）。

时间复杂度：
平均开销：O(nlogn)
最坏情况：O(n^2)
如果数组中的数字是有序的，分治算法将没有任何优化效果。
空间复杂度：O(n)最坏情况下递归需要O(n)的空间

*/

func largestRectangleArea2(heights []int) int {

	return calculateArea(heights, 0, len(heights)-1)
}

func calculateArea(heights []int, start, end int) int {
	if start > end {
		return 0
	}
	minindex := start
	for i := start; i <= end; i++ {
		if heights[minindex] > heights[i] {
			minindex = i
		}
	}

	return max(heights[minindex]*(end-start+1), (calculateArea(heights, start, minindex-1)),
		(calculateArea(heights, minindex+1, end)))
}

func max(i, j, k int) int {
	if i > j && i > k {
		return i
	} else if j > k {
		return j
	}
	return k
}

/*

执行用时 :12 ms, 在所有 Go 提交中击败了76.40%的用户
内存消耗 :6 MB, 在所有 Go 提交中击败了26.03%的用户

栈
在这种方法中，我们维护一个栈。一开始，我们把 -1 放进栈的顶部来表示开始。
初始化时，按照从左到右的顺序，我们不断将柱子的序号放进栈中，直到遇到相邻柱子呈下降关系，
也就是 a[i-1] > a[i] 。现在，我们开始将栈中的序号弹出，
直到遇到 stack[j] 满足a[stack[j]]≤a[i] 。每次我们弹出下标时，
我们用弹出元素作为高，形成的最大面积矩形的宽，是当前元素与 stack[top−1]
 之间的那些柱子。也就是当我们弹出 stack[top]stack[top] 时，
记当前元素在原数组中的下标为 i ，当前弹出元素为高的最大矩形面积为：
(i−stack[top−1]−1)×a[stack[top]].
更进一步，当我们到达数组的尾部时，我们将栈中剩余元素全部弹出栈。
在弹出每一个元素是，我们用下面的式子来求面积：(stack[top]−stack[top−1])×a[stack[top]]，
其中，stack[top]stack[top]表示刚刚被弹出的元素。
因此，我们可以通过每次比较新计算的矩形面积来获得最大的矩形面积。

时间复杂度：O(n)n 个数字每个会被压栈弹栈各一次。
空间复杂度： O(n)用来存放栈中元素。
*/

func largestRectangleArea(heights []int) int {
	max := 0
	stack := []int{0}
	heights = append([]int{-1}, heights...)
	heights = append(heights, 0)
	for i := 1; i < len(heights); i++ {
		if heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for {
			if heights[stack[len(stack)-1]] < heights[i] {
				stack = append(stack, i)
				break
			}
			top := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			area := top * (i - stack[len(stack)-1] - 1)
			if area > max {
				max = area
			}
		}
	}

	return max
}
