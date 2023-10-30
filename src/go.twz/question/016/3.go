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
	myList.Next.Next.Next.Next = &ListNode3{
		Value: 5,
	}
	myList.Next.Next.Next.Next.Next = &ListNode3{
		Value: 6,
	}

	newList := revListNode(myList)

	newList = reverseListNode3(newList, 4)

	newList = revListNode(newList)

	fmtListNode3(newList)
}

// 2023年10月29日11:01:54 这里直接将链表反转 先不管倒着反转的问题
func reverseListNode3(listNode *ListNode3, k int) *ListNode3 {
	// 我们的思路是 先反转 链表 再将按照key值进行一次反转 直接反转不在这里写 这里写按照顺序数量进行反转逻辑
	tmpListNode := listNode
	// 截取链表
	for i := 1; i < k; i++ {
		if tmpListNode == nil {
			break
		}
		tmpListNode = tmpListNode.Next
	}
	if tmpListNode == nil {
		return listNode
	}
	// 下一个需要截取的节点
	nextInterceptNode := tmpListNode.Next
	// 截断这个链表
	tmpListNode.Next = nil
	// 反转
	newListNode := revListNode(listNode)

	newNextListNode := reverseListNode3(nextInterceptNode, k)

	// 连接两个节点
	listNode.Next = newNextListNode
	return newListNode
}

// 这个是 直接 全部反转
// 将前一个节点的next改成前一个的前一个的改成nil 通过递归 让上一层的再把改成nil的下一个改成上上一个 以此循环
func revListNode(listNode *ListNode3) *ListNode3 {
	if listNode == nil || listNode.Next == nil {
		return listNode
	}
	result := revListNode(listNode.Next)
	listNode.Next.Next = listNode
	listNode.Next = nil
	return result
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
