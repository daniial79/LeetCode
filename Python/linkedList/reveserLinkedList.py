def reverse_list(self, head):
    current = head
    before = None
    after = None
    while current is not None:
        after = current.next
        current.next = before
        before = current
        current = after
    
    # returning before pointer as head of reversed linked list
    return before