package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := new(ListNode)
	tail := dummy

	ptr1, ptr2 := list1, list2

	for ptr1 != nil && ptr2 != nil {

		if ptr1.Val <= ptr2.Val {
			tail.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			tail.Next = ptr2
			ptr2 = ptr2.Next
		}

		tail = tail.Next
	}

	if ptr1 == nil {
		tail.Next = ptr2
		tail = tail.Next
	} else {
		tail.Next = ptr1
		tail = tail.Next
	}

	return dummy.Next
}
