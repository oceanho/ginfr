package linklist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	l := New()
	ele := "hello, world"
	for i := 0; i < 10000; i++ {
		l.Append(ele)
	}
	assert.Equal(t, 10000, l.Length())
	assert.Contains(t, l.All(), ele)
}

func TestLinkedList_All(t *testing.T) {
	l := New()
	l.Append(1, 2, 3, 4, 5, 6)
	assert.Equal(t, 6, l.Length())
	alls := l.All()
	assert.Equal(t, alls[0], 1)
	assert.Equal(t, alls[1], 2)
	assert.Equal(t, alls[l.Length()-1], 6)
}

func TestLinkedList_Slice(t *testing.T) {
	l := New()
	l.Append(1, 2, 3, 4, 5, 6)
	assert.Equal(t, 6, l.Length())
	alls := l.Slice(0, 1)
	assert.Equal(t, alls[0], 1)
	assert.Equal(t, len(alls), 1)

	alls = l.Slice(0, 6)
	assert.Equal(t, alls[0], 1)
	assert.Equal(t, alls[5], 6)
	assert.Equal(t, len(alls), 6)

	alls = l.Slice(2, 4)
	assert.Equal(t, alls[0], 3)
	assert.Equal(t, len(alls), 2)
}

func TestLinkedList_SliceWithStart(t *testing.T) {
	l := New()
	l.Append(1, 2, 3, 4, 5, 6)
	assert.Equal(t, 6, l.Length())
	alls := l.SliceWithStart(0)
	assert.Equal(t, alls[0], 1)
	assert.Equal(t, alls[5], 6)
	assert.Equal(t, len(alls), 6)

	alls = l.SliceWithStart(4)
	assert.Equal(t, alls[0], 5)
	assert.Equal(t, alls[1], 6)
	assert.Equal(t, len(alls), 2)
}

func TestLinkedList_SliceWithEnd(t *testing.T) {
	l := New()
	l.Append(1, 2, 3, 4, 5, 6)
	assert.Equal(t, 6, l.Length())
	alls := l.SliceWithEnd(6)
	assert.Equal(t, alls[0], 1)
	assert.Equal(t, alls[5], 6)
	assert.Equal(t, len(alls), 6)

	alls = l.SliceWithEnd(2)
	assert.Equal(t, alls[0], 1)
	assert.Equal(t, alls[1], 2)
	assert.Equal(t, len(alls), 2)
}

func TestLinkedList_Prepend(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Prepend(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	alls := l.All()
	assert.Contains(t, alls, ele)
	assert.Equal(t, alls[0], ele)
	for i := 1; i <= 5; i++ {
		assert.Equal(t, alls[i], i)
	}
}

func TestLinkedList_First(t *testing.T) {
	l := New()
	l.Append("hello, world", 1, 2, 3, 4, 5)
	assert.Equal(t, "hello, world", l.Head())
}

func TestLinkedList_Last(t *testing.T) {
	l := New()
	l.Append("hello, world", 1, 2, 3, 4, 5)
	assert.Equal(t, 5, l.Tail())
}

func TestLinkedList_JSON(t *testing.T) {
	l := New()
	l.Append("hello, world", struct {
		ID   int
		Name string
	}{1, "ocean"}, 1, 2, 3, 4, 5)
	bytes, err := l.JSON()
	assert.Nil(t, err)
	t.Logf("l.JSON, %v", string(bytes))
}

func TestLinkedList_Update(t *testing.T) {
	l := New()
	ele := "hello, world"
	eleNew := "hello, world, hahaha."
	l.Append(ele, 1, 2, 3, 4, 5)
	err := l.Update(ele, eleNew)
	assert.Nil(t, err)
	val := l.Find(func(values ...interface{}) bool {
		if s, ok := values[0].(string); ok && s == eleNew {
			return true
		}
		return false
	})
	assert.Nil(t, err)
	assert.Equal(t, eleNew, val)

	val = l.Find(func(values ...interface{}) bool {
		if s, ok := values[0].(string); ok && s == ele {
			return true
		}
		return false
	})
	assert.Nil(t, val)

	assert.NotContains(t, l.All(), ele)
	assert.Contains(t, l.All(), eleNew)
}

func TestLinkedList_UpdateWithIndex(t *testing.T) {
	l := New()
	ele := "hello, world"
	eleNew := "hello, world, hahaha."
	l.Append(ele, 1, 2, 3, 4, 5)
	err := l.UpdateWithIndex(0, eleNew)
	assert.Nil(t, err)
	val := l.Index(0)
	assert.Nil(t, err)
	assert.Equal(t, eleNew, val)

	err = l.UpdateWithIndex(1, eleNew)
	assert.Nil(t, err)
	val = l.Index(1)
	assert.Nil(t, err)
	assert.Equal(t, eleNew, val)

	err = l.UpdateWithIndex(l.Length()-1, eleNew)
	assert.Nil(t, err)
	val = l.Index(l.Length() - 1)
	assert.Nil(t, err)
	assert.Equal(t, eleNew, val)
}

func TestLinkedList_UpdateWithExpr(t *testing.T) {
	l := New()
	ele := "hello, world"
	eleNew := "hello, world, hahaha."
	l.Append(ele, 1, 2, 3, 4, 5)
	err := l.UpdateWithExpr(func(values ...interface{}) bool {
		s, o := values[0].(string)
		return o && s == ele
	}, func(oldValue ...interface{}) interface{} {
		return eleNew
	})
	assert.Nil(t, err)
	val := l.Index(0)
	assert.Nil(t, err)
	assert.Equal(t, eleNew, val)

	err = l.UpdateWithExpr(func(values ...interface{}) bool {
		s, o := values[0].(int)
		return o && s == 3
	}, func(oldValue ...interface{}) interface{} {
		return oldValue[0].(int) * 2
	})
	assert.Nil(t, err)
	val = l.Index(3)
	assert.Nil(t, err)
	assert.Equal(t, 6, val)

	err = l.UpdateWithIndex(l.Length()-1, eleNew)
	assert.Nil(t, err)
	val = l.Index(l.Length() - 1)
	assert.Nil(t, err)
	assert.Equal(t, eleNew, val)
}

func TestLinkedList_Insert(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	err := l.Insert(0, 100, 101, 102, 103)
	assert.Nil(t, err)
	for i := 0; i < 4; i++ {
		assert.Equal(t, l.Index(i), 100+i)
	}
	err = l.Insert(l.Length(), 200)
	assert.NotNil(t, err)

	err = l.Insert(l.Length()-1, 200, 201, 202, 203)
	assert.Nil(t, err)
	assert.True(t, l.Exists(ele))
	start := l.Length() - 4
	for i := start; i < start+4; i++ {
		assert.Equal(t, l.Index(i), 200-start+i)
	}
}

func TestLinkedList_Iter(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5, 100, 101, 102, 103)
	idx := 0
	l.Iter(func(values ...interface{}) {
		t.Logf("idx: %d, value: %v", idx, values)
		idx++
	})
}

func TestLinkedList_Exists(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.True(t, l.Exists(ele))
	for i := 1; i <= 5; i++ {
		assert.True(t, l.Exists(i))
	}
	assert.False(t, l.Exists(100))
}

func TestLinkedList_ExistsWithExpr(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.True(t, l.ExistsWithExpr(func(values ...interface{}) bool {
		s, o := values[0].(int)
		return o && s == 5
	}))
	for i := 1; i <= 5; i++ {
		assert.True(t, l.ExistsWithExpr(func(values ...interface{}) bool {
			s, o := values[0].(int)
			return o && s == i
		}))
	}
	assert.False(t, l.ExistsWithExpr(func(values ...interface{}) bool {
		s, o := values[0].(int)
		return o && s == 6
	}))
}

func TestLinkedList_Remove(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	_ = l.Remove(ele)
	assert.Equal(t, 5, l.Length())
	assert.NotContains(t, l.All(), ele)

	_ = l.Remove(4)
	assert.Equal(t, 4, l.Length())
	assert.NotContains(t, l.All(), 4)
}

func TestLinkedList_RemoveWithExpr(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	_ = l.RemoveWithExpr(func(values ...interface{}) bool {
		if s, o := values[0].(string); o && s == ele {
			return true
		}
		return false
	})
	assert.NotContains(t, l.All(), ele)

	_ = l.Remove(4)
	assert.Equal(t, 4, l.Length())
	assert.NotContains(t, l.All(), 4)
}

func TestLinkedList_RemoveWithIndex(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	err := l.RemoveWithIndex(0)
	assert.Equal(t, err, nil)
	assert.Equal(t, 5, l.Length())
	assert.NotContains(t, l.All(), ele)

	err = l.RemoveWithIndex(6)
	assert.Equal(t, err, LinkedListNotMatchAnyElementWithIndexError)

	err = l.RemoveWithIndex(4)
	assert.Equal(t, 4, l.Length())
	assert.NotContains(t, l.All(), 5)
}

func TestLinkedList_RemoveWithRanger(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	err := l.RemoveWithRanger(0, l.Length()-1)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(l.All()), l.Length())
}

func TestLinkedList_Reset(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	l.Reset()
	assert.Equal(t, 0, l.Length())
	assert.Equal(t, len(l.All()), 0)
}

func TestLinkedList_PopHead(t *testing.T) {
	l := New()
	ele := "hello, world"
	l.Append(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	str, ok := l.PopHead().(string)
	assert.Equal(t, ok, true)
	assert.Equal(t, str, ele)
	assert.Equal(t, 1, l.Head())
	assert.Equal(t, 5, l.Tail())
	assert.Equal(t, l.Length(), len(l.All()))
	assert.NotContains(t, l.All(), ele)
	length := l.Length()
	for i := 0; i < length; i++ {
		val, ok := l.PopHead().(int)
		assert.Equal(t, ok, true)
		assert.Equal(t, val, i+1)
	}
	assert.Equal(t, 0, l.Length())
	assert.Equal(t, 0, len(l.All()))

	l.Append(ele, 1, 2, 3, 4, 5)
	assert.Equal(t, 6, l.Length())
	str, ok = l.PopHead().(string)
	assert.Equal(t, ok, true)
	assert.Equal(t, str, ele)
	assert.Equal(t, l.Length(), len(l.All()))
}

func TestLinkedList_PopTail(t *testing.T) {
	l := New()
	l.Append(1, 2, 3, 4, 5)
	val, ok := l.PopTail().(int)
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 5)
	assert.Equal(t, 4, l.Tail())
	assert.Equal(t, l.Length(), len(l.All()))
	assert.NotContains(t, l.All(), 5)
	length := l.Length()
	for i := 0; i < length; i++ {
		val, ok := l.PopTail().(int)
		assert.Equal(t, ok, true)
		assert.Equal(t, val, length-i)
	}
	assert.Equal(t, 0, l.Length())
	assert.Equal(t, 0, len(l.All()))

	l.Append(1, 2, 3, 4, 5)
	assert.Equal(t, 5, l.Length())
	val, ok = l.PopTail().(int)
	assert.Equal(t, ok, true)
	assert.Equal(t, val, 5)
	assert.Equal(t, l.Length(), len(l.All()))
}

func BenchmarkLinkedList_Append(b *testing.B) {
	l := New()
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}
}

func BenchmarkGoArray_Append(b *testing.B) {
	vals := make([]int, 0, 32)
	for i := 0; i < b.N; i++ {
		vals = append(vals, i)
	}
}

func BenchmarkLinkedList_Prepend(b *testing.B) {
	l := New()
	for i := 0; i < b.N; i++ {
		l.Prepend(i)
	}
}

func BenchmarkGoArray_Prepend(b *testing.B) {
	vals := make([]int, 0, 32)
	for i := 0; i < b.N; i++ {
		tmp := make([]int, 0, len(vals))
		copy(tmp, vals)
		vals = append(vals, i)
		vals = append(vals, tmp...)
	}
}
