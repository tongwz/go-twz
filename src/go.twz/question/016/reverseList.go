package main

import "fmt"

type ListNode4 struct {
	Value int
	Next  *ListNode4
}

func main() {
	var myList = &ListNode4{
		Value: 1,
	}
	myList.Next = &ListNode4{
		Value: 2,
	}
	myList.Next.Next = &ListNode4{
		Value: 3,
	}
	myList.Next.Next.Next = &ListNode4{
		Value: 4,
	}
	myList.Next.Next.Next.Next = &ListNode4{
		Value: 5,
	}
	myList.Next.Next.Next.Next.Next = &ListNode4{
		Value: 6,
	}

	newList := reverseList1(myList)

	fmtListNode4(newList)

	newList = reverseList2(newList)

	fmtListNode4(newList)
}

func reverseList1(listNode *ListNode4) *ListNode4 {
	var pre *ListNode4
	var next *ListNode4
	for {
		if listNode == nil {
			break
		}
		next = listNode.Next
		listNode.Next = pre
		pre = listNode
		listNode = next
	}
	return pre
}

func reverseList2(listNode *ListNode4) *ListNode4 {
	if listNode == nil || listNode.Next == nil {
		return listNode
	}
	result := reverseList2(listNode.Next)
	listNode.Next.Next = listNode
	listNode.Next = nil
	return result
}

func fmtListNode4(tmpNode *ListNode4) {
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
