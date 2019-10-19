package linkedlist

import (
	"errors"
	"fmt"
)

// Todo come up with name
var ErrEmptyList = errors.New("error: list can't be empty")

type List struct {
	nodes []*Node
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

	var nodes []*Node
	root, rest := items[0], items[1:]
	temp := &Node{Val: root}
	nodes = append(nodes, temp)
	for _, item := range rest {
		nasd := &Node{Val: item, prev: temp}
		temp.next = nasd
		temp = nasd
		nodes = append(nodes, nasd)
	}
	return &List{nodes: nodes}
}

func (l *List) PushFront(item interface{})     {}
func (l *List) PushBack(item interface{})      {}
func (l *List) PopFront() (interface{}, error) { panic("") }
func (l *List) PopBack() (interface{}, error)  { panic("") }
func (l *List) First() *Node {
	if len(l.nodes) == 0 {
		return nil
	}
	return l.nodes[0]
}
func (l *List) Last() *Node {
	if len(l.nodes) == 0 {
		return nil
	}
	return l.nodes[len(l.nodes)-1]
}
func (l *List) Reverse() {}

func (l *List) String() string {
	var result string
	for _, node := range l.nodes {
		result += fmt.Sprintf("%v\n", node)
	}

	return result
}
