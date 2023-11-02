package main

import "fmt"

type listNode struct {
	Value int
	Next  *listNode
}

func reverseByKey(head *listNode, k int) *listNode {
	tmpHead := head
	for i := 0; i < k-1; i++ {
		if tmpHead == nil {
			break
		}
		tmpHead = tmpHead.Next
	}
	if tmpHead == nil {
		return head
	}

	nextHead := tmpHead.Next
	tmpHead.Next = nil

	newList := reverseList(head)

	leaveList := reverseByKey(nextHead, k)

	head.Next = leaveList

	return newList
}

func reverseList(head *listNode) *listNode {
	var pre *listNode
	var next *listNode
	for {
		if head == nil {
			break
		}
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverseList2(head *listNode) *listNode {
	if head.Next == nil {
		return head
	}
	res := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func reverseList3(head *listNode) *listNode {
	var newNext *listNode
	for head != nil {
		next := head.Next
		head.Next = newNext
		newNext = head
		head = next
	}
	return newNext
}

func main() {
	var myList = &listNode{
		Value: 1,
	}
	myList.Next = &listNode{
		Value: 2,
	}
	myList.Next.Next = &listNode{
		Value: 3,
	}
	myList.Next.Next.Next = &listNode{
		Value: 4,
	}
	myList.Next.Next.Next.Next = &listNode{
		Value: 5,
	}
	myList.Next.Next.Next.Next.Next = &listNode{
		Value: 6,
	}
	myList.Next.Next.Next.Next.Next.Next = &listNode{
		Value: 7,
	}
	myList.Next.Next.Next.Next.Next.Next.Next = &listNode{
		Value: 8,
	}

	fmtListNode(myList)

	newList := reverseList(myList)

	newList = reverseByKey(newList, 4)

	newList = reverseList(newList)

	fmtListNode(newList)

}

func fmtListNode(tmpNode *listNode) {
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
	}
}
