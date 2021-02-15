package types

type Iter func(values ...interface{})
type Filter func(values ...interface{}) bool
type SetValuer func(oldValue ...interface{}) interface{}
