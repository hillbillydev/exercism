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
	var (
		nodes    []*Node
		lastRecord = records[len(records)- 1]
		unique = make(map[int]bool)
	)

	// if the last records has an ID of 5 and the lenght is 4 in the slice
	// it means that the slice is not continuous.
	if lastRecord.ID > len(records) - 1 {
		return nil, errors.New("Error: non-continuous nodes")
	}

	root, records := records[0], records[1:]
	if root.ID != root.Parent {
		return nil, fmt.Errorf("Error: Root node has a parent ID of %d", root.Parent)
	}

	rootNode := &Node{ID: root.ID}
	nodes = append(nodes, rootNode)

	for _, r := range records {
		if r.ID <= r.Parent {
			return nil, fmt.Errorf("Error: didn't expect ID of %d to be less of equal to the parent of (%d)", r.ID, r.Parent)
		}

		if _, exists := unique[r.ID]; exists {
			return nil, fmt.Errorf("Error: ID of %d is not unique", r.ID)
		}
		unique[r.ID] = true

		temp := &Node{ID: r.ID}
		nodes = append(nodes, temp)
		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, temp)
	}

	return rootNode, nil
}

