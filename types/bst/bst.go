package bst

import (
	"github.com/oceanho/ginfr/types"
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
	parent *node
	left   *node
	right  *node
	data   *linklist.LinkedList
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

func newNode(value interface{}, parent, left, right *node) *node {
	list := linklist.New()
	if value != nil {
		list.Append(value)
	}
	return &node{
		parent: parent,
		left:   left,
		right:  right,
		data:   list,
	}
}

func (bst *BST) visit(value interface{}, parent, node *node) {
	if node == nil {
		node = newNode(value, parent, nil, nil)
		return
	}
	if node.data.Head() == nil {
		node.data.Append(value)
		return
	}

	if bst.comparer.Compare(node.data.Head(), value) < 0 {
		bst.visit(value, node, node.left)
	} else if bst.comparer.Compare(node.data.Head(), value) == 0 {
		node.data.Append(value)
	} else if bst.comparer.Compare(node.data.Head(), value) > 0 {
		bst.visit(value, node, node.right)
	}
}

func (bst *BST) insert(value interface{}) {
	if bst.root == nil {
		bst.root = newNode(value, nil, nil, nil)
		return
	}
	bst.visit(value, bst.root, nil)
}

func (bst *BST) bfs(node *node, handler types.BFSHandler) {
	if node == nil || node.data == nil || handler(node.data.All()) {
		return
	}
	vals := node.data.All()
	for len(vals) > 0 {
		tmps := vals[0:]
		for _, v := range tmps {
			if handler(v) {
				return
			}
		}
		vals = vals[:0]
		if node.left != nil {
			vals = append(vals, node.left)
		}
		if node.right != nil {
			vals = append(vals, node.right)
		}
	}
}

func (bst *BST) dfs(level int, node *node, handler types.DFSHandler) {
	if node == nil || node.data == nil || handler(level, node.data.All()) {
		return
	}
	bst.dfs(level+1, node.left, handler)
	bst.dfs(level+1, node.right, handler)
}

func (bst *BST) Insert(values ...interface{}) {
	for _, value := range values {
		bst.insert(value)
	}
}

// Pre-order
// Parent -> Left -> Right
func (bst *BST) PreOrderIter(iter types.Iter) {
	bst.preOrderIter(iter, bst.root)
}

// Pre-order Loop.
// Parent -> Left -> Right
func (bst *BST) preOrderIter(iter types.Iter, node *node) {
	if node == nil {
		return
	}
	values := node.data
	values.Iter(iter)
	bst.preOrderIter(iter, node.left)
	bst.preOrderIter(iter, node.right)
}

// Pre-order Loop.
// Parent -> Left -> Right
func (bst *BST) preOrderIterInternal(iter types.Filter, node *node) {

}

// In-order Loop.
// Left, Parent, Right
func (bst *BST) inOrderIter(iter types.Filter) {

}

// Post-order Loop.
// Right, Parent, Left
func (bst *BST) postOrderIter(iter types.Filter) {

}

func (bst *BST) PreOrderList(filter types.Filter) []interface{} {
	val := make([]interface{}, 0, 128)
	bst.PreOrderIter(func(values ...interface{}) {
		val = append(val, values[0])
	})
	return val
}
