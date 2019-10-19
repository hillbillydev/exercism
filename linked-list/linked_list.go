package linkedlist

import (
	"errors"
)

var ErrEmptyList = errors.New("error: list can't be empty")

type List struct {
	*Node
}

type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) First() *Node {
	if n == nil {
		return nil
	}
	last := n
	for {
		node := last.Prev()
		if node == nil {
			return last
		}
		last = node
	}
	return nil
}

func (n *Node) Last() *Node {
	if n == nil {
		return nil
	}

	last := n
	for {
		node := last.Next()
		if node == nil {
			return last
		}
		last = node
	}
	return nil
}

func NewList(items ...interface{}) *List {
	if len(items) < 1 {
		return &List{}
	}

	root, rest := items[0], items[1:]
	rootNode := &Node{Val: root}
	prev := rootNode
	for _, item := range rest {
		tmp := &Node{
			Val:  item,
			prev: prev,
		}

		prev.next = tmp
		prev = tmp
	}

	return &List{Node: rootNode}
}

func (l *List) PushFront(item interface{})     {}
func (l *List) PushBack(item interface{})      {}
func (l *List) PopFront() (interface{}, error) { panic("") }
func (l *List) PopBack() (interface{}, error)  { panic("") }

func (l *List) Reverse() {
	//	var nodes []*Node
	//
	//	last := l.Last()
	//	for {
	//		node := last.Prev()
	//		if node == nil {
	//			break
	//		}
	//		last = node
	//		nodes = append(nodes, node)
	//	}
}

func (l *List) String() string {
	//var result string
	//for _, node := range l.nodes {
	//	result += fmt.Sprintf("%v\n", node)
	//}

	return ""
}
