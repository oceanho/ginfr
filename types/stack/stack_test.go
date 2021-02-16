package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stk := New()
	for i := 0; i < 2; i++ {
		stk.Push(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		assert.Equal(t, 10, stk.Length())
		val := stk.Peek()
		assert.Equal(t, val, 10)
		assert.False(t, stk.IsEmpty())
		assert.Equal(t, 10, stk.Length())
		for i := 0; i < 10; i++ {
			val := stk.Pop()
			assert.Equal(t, val, 10-i)
			assert.Equal(t, 10-i-1, stk.Length())
		}
		assert.True(t, stk.IsEmpty())

		stk.Push(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		stk.Clear()
		assert.True(t, stk.IsEmpty())

		stk.Push(1)
		stk.Push("hello")
		v1, ok1 := stk.Peek().(string)
		assert.True(t, ok1)
		assert.Equal(t, "hello", v1)
		assert.Equal(t, 2, stk.Length())

		v1, ok1 = stk.Pop().(string)
		assert.True(t, ok1)
		assert.Equal(t, "hello", v1)
		assert.Equal(t, 1, stk.Length())

		v, ok := stk.Peek().(int)
		assert.True(t, ok)
		assert.Equal(t, 1, v)
		assert.Equal(t, 1, stk.Length())

		v2, ok2 := stk.Pop().(int)
		assert.True(t, ok2)
		assert.True(t, stk.IsEmpty())
		assert.Equal(t, 1, v2)
		assert.Equal(t, 0, stk.Length())
	}
}

func BenchmarkStack_Push(b *testing.B) {
	stk := New()
	for i := 0; i < b.N; i++ {
		stk.Push(i)
	}
}
