package types

type Iter func(values ...interface{})
type Filter func(values ...interface{}) bool
type SetValuer func(oldValue ...interface{}) interface{}

func Truer(values ...interface{}) bool {
	return true
}

func Falser(values ...interface{}) bool {
	return false
}
