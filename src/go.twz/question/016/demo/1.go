package main

import "fmt"

type listNode struct {
	Value int
	Next  *listNode
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

	fmtListNode(myList)

	newList := reverseList(myList)

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
