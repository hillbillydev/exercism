package tree

import (
	"fmt"
	"sort"
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

	root := records[0]
	if root.ID != root.Parent {
		return nil, fmt.Errorf("Error: Root node doesn't have the same ID(%q) as ParentID(%q)", root.ID, root.Parent)
	}

	var (
		nodes    = make([]*Node, len(records))
		rootNode = &Node{ID: records[0].ID}
	)
	nodes = append(nodes, rootNode)

	for _, r := range records {
		if r.ID <= r.Parent {
			return nil, fmt.Errorf("Error: didn't expect ID of %q to be less of equal to the parent of (%q)", r.ID, r.Parent)
		}
		temp := &Node{ID: r.ID}
		nodes = append(nodes, temp)
		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, temp)
	}

	return rootNode, nil
}
