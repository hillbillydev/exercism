package tree

import (
	"fmt"
	"sort"
	"errors"
)

// Record represent an simple unstructured data structure
type Record struct {
	ID, Parent int
}

// Node represents a single node in a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build takes records and transforms it to a tree.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	root, records := records[0], records[1:]
	if root.ID != root.Parent {
		return nil, fmt.Errorf("Error: Root node doesn't have the same ID(%d) as ParentID(%d)", root.ID, root.Parent)
	}

	var (
		nodes    []*Node
		rootNode = &Node{ID: root.ID}
	)

	nodes = append(nodes, rootNode)

	var temp *Node
	for _, r := range records {
		if r.ID <= r.Parent {
			return nil, fmt.Errorf("Error: didn't expect ID of %d to be less of equal to the parent of (%d)", r.ID, r.Parent)
		}

		temp = &Node{ID: r.ID}
		nodes = append(nodes, temp)
		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, temp)
	}

	m := make(map[int]bool)
	for _, n := range nodes {
		if _, exists := m[n.ID]; exists {
			return nil, fmt.Errorf("Error: ID of %d is not unique.", n.ID)
		}
		m[n.ID] = true
	}

	if nodes[len(nodes)- 1].ID > len(nodes) - 1 {
		return nil, errors.New("Error: non-continuous nodes")
	}

	return rootNode, nil
}
