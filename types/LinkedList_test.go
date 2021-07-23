package types

import (
	"fmt"
	"testing"
)

func TestListNode_Reverse(t *testing.T) {
	list := newListNode(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})

	list.Reverse().List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})

	list.Reverse().List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})
}

func TestListNode_ReverseV2(t *testing.T) {
	list := newListNode(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})

	list.ReverseV2().List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})

	list.ReverseV2().List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})
}

func TestListNode_ReverseV3(t *testing.T) {
	list := newListNode(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})
	l2 := reverse(list)
	l2.List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})
	reverse(l2).List(func(value interface{}) bool {
		fmt.Printf("%v\n", value)
		return false
	})
}

func BenchmarkListNode_ReverseV2(b *testing.B) {
	list := newListNode(1)
	list.Insert(2).Insert(3).Insert(4).Insert(5)
	for i := 0; i < b.N; i++ {
		list.ReverseV2()
	}
}

func BenchmarkListNode_ReverseV3(b *testing.B) {
	list := newListNode(1)
	list.Insert(2).Insert(3).Insert(4).Insert(5)
	for i := 0; i < b.N; i++ {
		reverse(list)
	}
}

func BenchmarkListNode_Reverse(b *testing.B) {
	list := newListNode(1)
	list.Insert(2).Insert(3).Insert(4).Insert(5)
	for i := 0; i < b.N; i++ {
		list.Reverse()
	}
}

func BenchmarkListNode_List(b *testing.B) {
	list := newListNode(1)
	for i := 0; i < 10000; i++ {
		list.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		list.List(func(value interface{}) bool {
			return false
		})
	}
}
