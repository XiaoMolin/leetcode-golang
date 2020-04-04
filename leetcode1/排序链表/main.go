// 排序链表 project main.go
package main

import (
	"fmt"
)

func main() {
	head := ListNode{4, nil}
	head.Next = &ListNode{2, nil}
	a := ListNode{1, nil}
	head.Next.Next = &a
	b := ListNode{3, nil}
	a.Next = &b
	fmt.Println(sortList(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
执行用时 :12 ms, 在所有 Go 提交中击败了92.43%的用户
内存消耗 :5.1 MB, 在所有 Go 提交中击败了66.51%的用户

时间复杂度O(logN)
空间复杂度O(1)

归并排序
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head //初始化慢指针和快指针
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next      //慢指针走一步
		fast = fast.Next.Next //快指针走两步
	}
	right := sortList(slow.Next)            //递归两分左链表
	slow.Next = nil                         //将前半个链表断开
	left := sortList(head)                  //递归两分右链表
	return mergeTwoSortedLists(left, right) //合并排序后的两端链表
}

func mergeTwoSortedLists(l, r *ListNode) *ListNode {
	res := new(ListNode)
	p := res
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l != nil {
		p.Next = l
	}
	if r != nil {
		p.Next = r
	}
	return res.Next
}
