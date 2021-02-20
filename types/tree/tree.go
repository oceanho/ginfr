package tree

type node struct {
	data     interface{}
	children []*node
}

type head struct {
	len  int
	cap  int
	deep int
}

type BaseTree struct {
	head *head
	root *node
}
