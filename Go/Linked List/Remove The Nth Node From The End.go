package sll

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	scout, prev := head, head

	for i := 0; i < n; i++ {
		scout = scout.Next
	}

	if scout == nil {
		return head.Next
	}

	for scout.Next != nil {
		scout = scout.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next

	return head

}
