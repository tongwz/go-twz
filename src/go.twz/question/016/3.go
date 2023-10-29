package main

import "fmt"

type ListNode3 struct {
	Value int
	Next  *ListNode3
}

func main() {
	var myList = &ListNode3{
		Value: 1,
	}
	myList.Next = &ListNode3{
		Value: 2,
	}
	myList.Next.Next = &ListNode3{
		Value: 3,
	}
	myList.Next.Next.Next = &ListNode3{
		Value: 4,
	}

	newList := revListNode(myList)

	fmtListNode3(newList)
}

// 2023年10月29日11:01:54 这里直接将链表反转 先不管倒着反转的问题
// func reverseListNode3(listNode *ListNode3, key int) *ListNode3 {
// 	// 拿到链表长度
// 	var lenList = LenListNode(listNode)
// 	if lenList < key {
// 		return listNode
// 	}
// 	times := lenList / key
// 	for i := 1; i <= times; i++ {
//
// 	}
// }

// 完全反正链表 每一个只要跟它的下一个进反转就行了,第一个变最后一个并且Next没有值
func revListNode(listNode *ListNode3) *ListNode3 {
	current := listNode      // tmpList = A指针
	if current.Next != nil { // 如果B指针 ！= nil
		tmpA := current
		tmpB := current.Next
		// tmpC := current.Next.Next // tmp = C指针
		current = tmpB // tmpList = B指针
		current.Next = tmpA
	}
	return nil
}

func LenListNode(listNode *ListNode3) (lenList int) {
	tmpListNode := listNode
	for {
		if tmpListNode != nil {
			lenList++
		} else {
			break
		}
		tmpListNode = listNode.Next
	}
	return lenList
}

func fmtListNode3(tmpNode *ListNode3) {
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
