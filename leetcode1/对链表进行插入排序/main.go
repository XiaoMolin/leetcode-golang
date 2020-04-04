// 对链表进行插入排序 project main.go
package main

import (
	"fmt"
)

func main() {
	head := ListNode{4, nil}
	head1 := ListNode{2, nil}
	head2 := ListNode{1, nil}
	head3 := ListNode{3, nil}
	head.Next = &head1
	head1.Next = &head2
	head2.Next = &head3
	fmt.Println(insertionSortList(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
执行用时 :24 ms, 在所有 Go 提交中击败了56.50%的用户
内存消耗 :3.3 MB, 在所有 Go 提交中击败了62.35%的用户
*/
func insertionSortList1(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		curr := dummy
		next := head.Next
		for curr.Next != nil && curr.Next.Val < head.Val {
			curr = curr.Next
		}

		head.Next = curr.Next
		curr.Next = head
		head = next
	}
	return dummy.Next
}

/*
执行用时 :4 ms, 在所有 Go 提交中击败了99.00%的用户
内存消耗 :3.3 MB, 在所有 Go 提交中击败了91.76%的用户
*/
func insertionSortList(head *ListNode) *ListNode {
	headpre := &ListNode{Next: head}
	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val < p.Val {
			cur = p
			continue
		}
		cur.Next = p.Next
		pre, next := headpre, headpre.Next
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		pre.Next = p
		p.Next = next
	}
	return headpre.Next
}
