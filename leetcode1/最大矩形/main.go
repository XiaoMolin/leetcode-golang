// 最大矩形 project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	fmt.Println(maximalRectangle(matrix))
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了82.42%的用户
内存消耗 :4.4 MB, 在所有 Go 提交中击败了35.21%的用户
运用84题柱状图中最大的矩形方法
*/
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	//保存最终结果
	max := 0
	//行数，列数
	m, n := len(matrix), len(matrix[0])
	//高度数组（统计每一行中1的高度）
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//每一行去找1的高度
			//如果不是‘1’，则将当前高度置为0
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				//是‘1’，则将高度加1
				height[j] = height[j] + 1
			}
		}
		//更新最大矩形的面积
		max = int(math.Max(float64(max), float64(largestRectangleArea(height))))
	}
	return max
}

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
