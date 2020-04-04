// 两棵二叉搜索树中的所有元素 project main.go
package main

import (
	"fmt"
)

func main() {
	root1 := TreeNode{Val: 2}
	root1.Left = &TreeNode{1, nil, nil}
	root1.Right = &TreeNode{4, nil, nil}
	root2 := TreeNode{Val: 1}
	root2.Left = &TreeNode{0, nil, nil}
	root2.Right = &TreeNode{3, nil, nil}
	fmt.Println(getAllElements(&root1, &root2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	a1, a2 := []int{}, []int{}
	ldf(root1, &a1)
	ldf(root2, &a2)
	return mergeSort(a1, a2)
}

// 中序遍历
func ldf(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		ldf(root.Left, a)
	}
	*a = append(*a, root.Val)
	fmt.Println(a)
	if root.Right != nil {
		ldf(root.Right, a)
	}
}

// 归并排序
func mergeSort(a1, a2 []int) []int {
	arr := []int{}
	i1, i2 := 0, 0
	for i1 < len(a1) && i2 < len(a2) {
		if a1[i1] <= a2[i2] {
			arr = append(arr, a1[i1])
			i1++
		} else {
			arr = append(arr, a2[i2])
			i2++
		}
	}
	if i1 < len(a1) {
		arr = append(arr, a1[i1:]...)
	}
	if i2 < len(a2) {
		arr = append(arr, a2[i2:]...)
	}
	return arr
}
