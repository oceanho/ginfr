package types

type ListNode struct {
	data interface{}
	next *ListNode
}

type LoopHandler func(value interface{}) bool

func (list *ListNode) newNextListNode(data interface{}) *ListNode {
	list.next = newListNode(data)
	return list
}

func (list *ListNode) setNextListNode(next *ListNode) {
	list.next = next
}

func newListNode(data interface{}) *ListNode {
	return &ListNode{
		data: data,
	}
}

func (list *ListNode) Insert(data interface{}) *ListNode {
	lastNode := list
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newListNode(data)
	return list
}

func (list *ListNode) Update(data interface{}) *ListNode {
	list.next = newListNode(data)
	return list
}

func (list *ListNode) Delete(data interface{}) *ListNode {
	list.next = newListNode(data)
	return list
}

func (list *ListNode) List(handler LoopHandler) {
	cur := list
	for cur != nil {
		if handler(cur.data) {
			break
		}
		cur = cur.next
	}
}

func (list *ListNode) ToArray(value chan interface{}) {
	cur := list
	for cur != nil {
		value <- cur.data
		cur = cur.next
	}
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.next = cur, cur.next, pre
	}
	return pre
}

func (list *ListNode) ReverseV2() *ListNode {
	return reverse(list)
}

func (list *ListNode) Reverse() *ListNode {
	if list == nil {
		return nil
	}
	next := list
	var nodes []*ListNode
	for next != nil {
		nodes = append(nodes, next)
		next = next.next
	}
	if nodes == nil {
		return nil
	}
	mi := len(nodes) - 1
	if mi < 0 {
		return nil
	}
	head := newListNode(nodes[mi].data)
	for i := len(nodes) - 2; i >= 0; i-- {
		head.Insert(nodes[i].data)
	}
	return head
}
