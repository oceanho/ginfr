package stack

import "github.com/oceanho/ginfr/types/linklist"

type Stack struct {
	l *linklist.LinkedList
}

func New() *Stack {
	return &Stack{
		l: linklist.New(),
	}
}

func (stk *Stack) IsEmpty() bool {
	return stk.l.IsEmpty()
}

func (stk *Stack) Length() int {
	return stk.l.Length()
}

func (stk *Stack) Push(values ...interface{}) {
	stk.l.Prepend(values...)
}

func (stk *Stack) Pop() interface{} {
	return stk.l.PopTail()
}

func (stk *Stack) Peek() interface{} {
	return stk.l.Tail()
}

func (stk *Stack) Clear() {
	_ = stk.l.RemoveWithRanger(0, stk.l.Length()-1)
}
