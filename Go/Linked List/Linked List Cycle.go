package sll

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	tortoise, hare := head, head

	for hare.Next != nil && hare.Next.Next != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next

		if hare == tortoise {
			return true
		}
	}

	return false
}
