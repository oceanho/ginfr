package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST(t *testing.T) {
	var l, r interface{}
	l = struct {
		ID   int
		Name string
	}{1, "ocean"}
	r = struct {
		ID   int
		Name string
	}{1, "ocean"}
	assert.True(t, l == r)
}

type myComparer struct {
}

func (m myComparer) Compare(left, right interface{}) int {
	return left.(int) - right.(int)
}

func TestBST_Insert(t *testing.T) {
	bst := New(&myComparer{})
	bst.Insert(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
}
