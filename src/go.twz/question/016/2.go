package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	var myList = &ListNode{
		Value: 1,
	}
	myList.Next = &ListNode{
		Value: 2,
	}
	myList.Next.Next = &ListNode{
		Value: 3,
	}
	myList.Next.Next.Next = &ListNode{
		Value: 4,
	}
	myList.Next.Next.Next.Next = &ListNode{
		Value: 5,
	}
	myList.Next.Next.Next.Next.Next = &ListNode{
		Value: 6,
	}
	myList.Next.Next.Next.Next.Next.Next = &ListNode{
		Value: 7,
	}
	myList.Next.Next.Next.Next.Next.Next.Next = &ListNode{
		Value: 8,
	}

	fmtListNode(myList)

	newList := reverseKGroup(myList, 3)

	fmtListNode(newList)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 计算链表长度
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}

	// 创建一个虚拟头节点作为结果链表的头部
	dummy := &ListNode{Next: head}
	prevGroupTail := dummy

	for length >= k {
		// 找到当前组的头节点和尾节点
		groupHead := prevGroupTail.Next
		groupTail := groupHead
		for i := 1; i < k; i++ {
			groupTail = groupTail.Next
		}

		// 断开当前组和下一组的连接
		nextGroupHead := groupTail.Next
		groupTail.Next = nil

		// 反转当前组
		prev, curr := nextGroupHead, groupHead
		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// 将反转后的组重新连接到结果链表中
		prevGroupTail.Next = groupTail
		groupHead.Next = nextGroupHead

		// 更新 prevGroupTail 和 length 的值
		prevGroupTail = groupHead
		length -= k
	}

	return dummy.Next
}

func fmtListNode(tmpNode *ListNode) {
	for {
		if tmpNode != nil {
			fmt.Printf("%d,", tmpNode.Value)
		} else {
			break
		}
		tmpNode = tmpNode.Next
		if tmpNode == nil {
			fmt.Println("")
		}
		// fmt.Printf("%#v \n", *tmpNode.Next)
		// time.Sleep(time.Millisecond * 300)
	}
}
