package linklist

import "github.com/oceanho/ginfr/types"

func updater(expr types.Filter, valuer types.SetValuer, values ...interface{}) bool {
	if expr(values[2]) {
		values[0].(*node).val = valuer(values[2])
		return true
	}
	return false
}

func remover(expr types.Filter, values ...interface{}) bool {
	if expr(values[2]) {
		node := values[0].(*node)
		node.next = node.next.next
		return true
	}
	return false
}
