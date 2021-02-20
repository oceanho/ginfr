package bst

import (
	"github.com/oceanho/ginfr/types/linklist"
)

// Valuer define a bst value compare feature object.
type Valuer interface {
	// Compare define a API that for compare to other one
	// Return rules are
	//	-1, if current > other
	//	0, if current == other
	//	1, if current < other
	Compare(other Valuer) int
}

type head struct {
	len  int
	cap  int
	deep int
}

type node struct {
	prev  *node
	left  *node
	right *node
	data  *linklist.LinkedList
}

// BST define a data structures of Binary Search Tree.
type BST struct {
	head     *head
	root     *node
	comparer Comparer
}

type Comparer interface {
	// Compare define a API that for compare to other one
	// Return rules are
	//	-1, if left > right
	//	0, if left == right
	//	1, if left < right
	Compare(left, right interface{}) int
}

func New(comparer Comparer) *BST {
	return &BST{
		head: &head{
			len:  0,
			cap:  0,
			deep: 0,
		},
		comparer: comparer,
		root:     newNode(nil, nil, nil, nil),
	}
}

func newNode(value interface{}, prev, left, right *node) *node {
	list := linklist.New()
	if value != nil {
		list.Append(value)
	}
	return &node{
		prev:  prev,
		left:  left,
		right: right,
		data:  list,
	}
}

func (bst *BST) left(value interface{}, parent *node) {
	node := parent.left
	if node == nil {
		parent.left = newNode(value, parent, nil, nil)
		return
	}
	if node.data.Head() == nil {
		node.data.Append(value)
		return
	}
	if bst.comparer.Compare(parent.data.Head(), value) < 0 {
		bst.left(value, node)
	} else if bst.comparer.Compare(parent.data.Head(), value) == 0 {
		node.data.Append(value)
	} else if bst.comparer.Compare(parent.data.Head(), value) > 0 {
		bst.right(value, node)
	}
}

func (bst *BST) right(value interface{}, parent *node) {
	node := parent.right
	if node == nil {
		parent.right = newNode(value, parent, nil, nil)
		return
	}
	if node.data.Head() == nil {
		node.data.Append(value)
		return
	}
	if bst.comparer.Compare(parent.data.Head(), value) < 0 {
		bst.left(value, node)
	} else if bst.comparer.Compare(parent.data.Head(), value) == 0 {
		node.data.Append(value)
	} else if bst.comparer.Compare(parent.data.Head(), value) > 0 {
		bst.right(value, node)
	}
}

func (bst *BST) node(value interface{}) {
	node := bst.root
	if node.data.Head() == nil {
		node.data.Append(value)
		return
	}
	if bst.comparer.Compare(node.data.Head(), value) < 0 {
		bst.left(value, node)
	} else if bst.comparer.Compare(node.data.Head(), value) == 0 {
		node.data.Append(value)
	} else if bst.comparer.Compare(node.data.Head(), value) > 0 {
		bst.right(value, node)
	}
}

func (bst *BST) insert(value interface{}) {
	bst.node(value)
}

func (bst *BST) bfs() *node {
	return nil
}

func (bst *BST) dfs() *node {
	return nil
}

func (bst *BST) Insert(values ...interface{}) {
	for _, value := range values {
		bst.insert(value)
	}
}
