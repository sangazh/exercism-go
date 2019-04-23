package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// a refactored tree building algorithm
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	//verify
	if err := verify(records); err != nil {
		return nil, err
	}

	//sort
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)
	for _, r := range records {
		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID != r.Parent {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}

	return nodes[0], nil
}

// verify if records valid
func verify(records []Record) error {
	recordMap := make(map[int]Record)
	for _, record := range records {
		recordMap[record.ID] = record

		if record.ID >= len(records) {
			return errors.New("node id is larger than or equal to length")
		}

		if record.ID == 0 && record.Parent != 0 {
			return errors.New("root id != parent")
		}
		if record.ID > 0 && record.Parent >= record.ID {
			return errors.New("node parent should be smaller than id")
		}
	}

	if _, ok := recordMap[0]; !ok {
		return errors.New("no root node")
	}

	if len(recordMap) != len(records) {
		return errors.New("duplicate")
	}

	return nil
}
