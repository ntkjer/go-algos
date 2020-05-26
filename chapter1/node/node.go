package node

/*
Node: item : interface {} and points to null by default
*/
type Node struct {
	item interface{}
	next *Node
}
