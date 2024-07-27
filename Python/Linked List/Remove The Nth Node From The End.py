class NodeList:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next
        
    def __str__(self) -> str:
        return f"{self.value}"
        
        
def remove_nth_from_end(head: NodeList, n: int) -> NodeList:
    
    if head is None or head.next is None:
        return head
    
    
    scout, prev = head, head
    
    for _ in range(n):
        scout = scout.next
        
    if scout is None or scout.next is None:
        return head.next

    prev.next = prev.next.next
    
    return prev

    
def print_list(head: NodeList):
    current = head
    
    while current is not None:
        print(f"{current.value} => ", end="")    
        current = current.next

    print("null")



def main():
    head = NodeList(1, NodeList(2, NodeList(3, NodeList(4, None))))
    
    
    print_list(head)
    
    remove_nth_from_end(head, 3)
    
    print_list(head)
    
    
if __name__ == '__main__':
    main()
    
    
    