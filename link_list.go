package sensitive

// LinkList ...
type LinkList struct {
	head  *listNode
	tail  *listNode
	count int64
}

// Push appends a node
func (list *LinkList) Push(v interface{}) {
	node := &listNode{
		Value: v,
	}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.Next = node

	}
	list.tail = node
	list.count++
}

// Pop returns the value of the first node
func (list *LinkList) Pop() interface{} {
	if list.Empty() {
		return nil
	}

	n := list.head
	list.head = n.Next
	list.count--
	return n.Value
}

// Empty returns true if there is none node
func (list *LinkList) Empty() bool {
	return list.count == 0
}

type listNode struct {
	Value interface{}
	Next  *listNode
}
