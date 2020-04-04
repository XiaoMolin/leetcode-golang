// 两数相加 project main.go
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Add(head *ListNode, data int) {
	point := head //临时指针
	for point.Next != nil {
		point = point.Next //移位
	}
	var node ListNode  //新节点
	point.Next = &node //赋值
	node.Val = data

}
func main() {
	var l1 ListNode = ListNode{Val: 0, Next: nil}
	for i := 0; i < 1; i++ {
		var str int
		fmt.Println("111111111111111输入")
		fmt.Scanln(&str)
		Add(&l1, str)
	}
	var l2 ListNode = ListNode{Val: 0, Next: nil}
	for i := 0; i < 10; i++ {
		var str int
		fmt.Println("2222222222222输入")
		fmt.Scanln(&str)
		Add(&l2, str)
	}
	addTwoNumbers(&l1, &l2)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := l1
	list2 := l2
	list1 = list1.Next
	list2 = list2.Next
	var l ListNode = ListNode{Val: 0, Next: nil}
	point := &l
	var carry int = 0
	for list1 != nil || list2 != nil {
		var x, y int
		if list1 != nil {
			x = list1.Val
		} else {
			x = 0
		}
		if list2 != nil {
			y = list2.Val
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		var node ListNode  //新节点
		point.Next = &node //赋值
		node.Val = sum % 10
		point = point.Next
		if list1 != nil {
			list1 = list1.Next
		}
		if list2 != nil {
			list2 = list2.Next
		}
	}
	if carry > 0 {
		var node ListNode  //新节点
		point.Next = &node //赋值
		node.Val = carry
		point = point.Next
	}
	pointt := &l //临时指针
	pointt = pointt.Next
	fmt.Println("bbbbbbbbbbbbbbbb")
	for pointt != nil {
		fmt.Println(pointt.Val)
		pointt = pointt.Next //移位
	}
	return pointt
}
