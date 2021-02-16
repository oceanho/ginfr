package linklist

import (
	"encoding/json"
	"fmt"
	"github.com/oceanho/ginfr/types"
)

type node struct {
	val  interface{}
	next *node
}

type head struct {
	len  int
	tail *node
}

type LinkedList struct {
	head *node
}

var (
	LinkedListNotMatchAnyElementError          = fmt.Errorf("linked list has no any match element")
	LinkedListNotMatchAnyElementWithIndexError = fmt.Errorf("linked list has no any match element with index")
)

func New() *LinkedList {
	return &LinkedList{
		head: newNode(&head{
			len:  0,
			tail: nil,
		}, nil),
	}
}

func newNode(val interface{}, next *node) *node {
	return &node{
		val:  val,
		next: next,
	}
}

func (l *LinkedList) sizeInc() {
	l.head.val.(*head).len += 1
}

func (l *LinkedList) sizeDec() {
	h := l.head.val.(*head)
	h.len -= 1
	if h.len < 0 {
		h.len = 0
	}
}

func (l *LinkedList) adjustTail(tail *node) {
	if tail != nil && tail == l.head {
		l.head.val.(*head).tail = nil
	} else {
		l.head.val.(*head).tail = tail
	}
}

// iter callback as (prev, curr, curr.val)
func (l *LinkedList) iter(iter types.Filter) {
	if l.IsEmpty() {
		return
	}
	stop := false
	prev := l.head
	curr := prev.next
	for curr != nil && !stop {
		stop = iter(prev, curr, curr.val)
		prev = curr
		curr = curr.next
	}
}

func (l *LinkedList) IsEmpty() bool {
	return l.head.next == nil
}

func (l *LinkedList) Length() int {
	return l.head.val.(*head).len
}

func (l *LinkedList) Iter(iter types.Iter) {
	l.iter(func(values ...interface{}) bool {
		iter(values[2])
		return false
	})
}

func (l *LinkedList) Append(values ...interface{}) {
	tail := l.head.val.(*head).tail
	if tail == nil {
		tail = newNode(values[0], nil)
		l.head.next = tail
		l.sizeInc()
		for _, el := range values[1:] {
			tail.next = newNode(el, nil)
			tail = tail.next
			l.sizeInc()
		}
	} else {
		for _, el := range values {
			tail.next = newNode(el, nil)
			tail = tail.next
			l.sizeInc()
		}
	}
	l.adjustTail(tail)
}

func (l *LinkedList) Prepend(values ...interface{}) {
	cur := l.head
	oldNext := l.head.next
	for _, el := range values {
		cur.next = newNode(el, nil)
		cur = cur.next
		l.sizeInc()
	}
	cur.next = oldNext
	if oldNext == nil {
		l.adjustTail(cur)
	}
}

func (l *LinkedList) Update(oldValue, newValue interface{}) error {
	return l.UpdateWithExpr(func(values ...interface{}) bool {
		return values[0] == oldValue
	}, func(oldValue ...interface{}) interface{} {
		return newValue
	})
}

func (l *LinkedList) UpdateWithIndex(idx int, newValue interface{}) error {
	index := 0
	return l.UpdateWithExpr(func(values ...interface{}) bool {
		if index == idx {
			return true
		}
		index++
		return false
	}, func(oldValue ...interface{}) interface{} {
		return newValue
	})
}

func (l *LinkedList) UpdateWithExpr(expr types.Filter, valuer types.SetValuer) error {
	err := LinkedListNotMatchAnyElementError
	l.iter(func(values ...interface{}) bool {
		if updater(expr, valuer, values...) {
			err = nil
			return true
		}
		return false
	})
	return err
}

func (l *LinkedList) Remove(value interface{}) error {
	return l.RemoveWithExpr(func(values ...interface{}) bool {
		return values[0] == value
	})
}

func (l *LinkedList) RemoveWithExpr(expr types.Filter) error {
	err := LinkedListNotMatchAnyElementError
	l.iter(func(values ...interface{}) bool {
		if remover(expr, values...) {
			l.sizeDec()
			return true
		}
		return false
	})
	return err
}

func (l *LinkedList) RemoveWithIndex(idx int) error {
	index := 0
	err := LinkedListNotMatchAnyElementWithIndexError
	l.iter(func(values ...interface{}) bool {
		if idx == index {
			err = nil
			prev := values[0].(*node)
			if prev.next == nil {
				l.adjustTail(prev)
			} else {
				prev.next = prev.next.next
			}
			l.sizeDec()
			return true
		}
		index++
		return false
	})
	return err
}

func (l *LinkedList) RemoveWithRanger(startIdx, endIdx int) error {
	index := 0
	err := LinkedListNotMatchAnyElementWithIndexError
	var left, right *node
	l.iter(func(values ...interface{}) bool {
		if startIdx == index {
			// left node should be are prev node object.
			left = values[0].(*node)
		} else if endIdx == index {
			// right node should be are current's next node object.
			right = values[1].(*node).next
		}
		index++
		return left != nil && (index == endIdx || right != nil)
	})
	if left != nil && (index == endIdx || right != nil) {
		err = nil
		left.next = right
		if right == nil {
			l.adjustTail(left)
		}
		l.head.val.(*head).len -= (endIdx - startIdx) + 1
	}
	return err
}

func (l *LinkedList) Reset() {
	h := l.head.val.(*head)
	h.len = 0
	h.tail = nil
	l.head.val = h
	l.head.next = nil
}

func (l *LinkedList) Find(expr types.Filter) interface{} {
	var val interface{} = nil
	l.iter(func(values ...interface{}) bool {
		if expr(values[2]) {
			val = values[2]
			return true
		}
		return false
	})
	return val
}

func (l *LinkedList) Head() interface{} {
	if l.head.next == nil {
		return nil
	}
	return l.head.next.val
}

func (l *LinkedList) Tail() interface{} {
	if tail := l.head.val.(*head).tail; tail != nil {
		return tail.val
	}
	return nil
}

func (l *LinkedList) PopHead() interface{} {
	head := l.head.next
	if head == nil {
		return nil
	}
	if head.next == nil {
		l.adjustTail(nil)
	}
	// Remove the first of top element.
	l.head.next = head.next
	l.sizeDec()
	return head.val
}

func (l *LinkedList) PopTail() interface{} {
	var prev, val interface{}
	l.iter(func(values ...interface{}) bool {
		if values[1].(*node).next == nil {
			prev, _, val = values[0], values[1], values[2]
			return true
		}
		return false
	})
	// Remove the last element.
	prev.(*node).next = nil
	l.adjustTail(prev.(*node))
	l.sizeDec()
	return val
}

func (l *LinkedList) All() []interface{} {
	return l.AllWithExpr(types.Truer)
}

func (l *LinkedList) Index(idx int) interface{} {
	index := 0
	var val interface{} = nil
	l.iter(func(values ...interface{}) bool {
		if index == idx {
			val = values[2]
			return true
		}
		index++
		return false
	})
	return val
}

func (l *LinkedList) Slice(start, end int) []interface{} {
	index := 0
	var returns = make([]interface{}, 0, 32)
	l.iter(func(values ...interface{}) bool {
		if index >= start && index < end {
			returns = append(returns, values[2])
		}
		index++
		return index >= end
	})
	return returns
}

func (l *LinkedList) SliceWithStart(start int) []interface{} {
	return l.Slice(start, l.Length())
}

func (l *LinkedList) SliceWithEnd(end int) []interface{} {
	return l.Slice(0, end)
}

func (l *LinkedList) AllWithExpr(expr types.Filter) []interface{} {
	var returns = make([]interface{}, 0, 32)
	l.iter(func(values ...interface{}) bool {
		if expr(values[2]) {
			returns = append(returns, values[2])
		}
		return false
	})
	return returns
}

func (l *LinkedList) JSON() ([]byte, error) {
	return json.Marshal(l.AllWithExpr(types.Truer))
}
