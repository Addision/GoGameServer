package main

import (
	"fmt"
	. "utils/linklist"
)

func main() {
	skipList := NewSkipList(12)

	skipList.Insert(1, 1)
	skipList.Insert(2, 2)
	skipList.Insert(2, 3)
	skipList.Insert(3, 3)
	skipList.Insert(5, 3)
	skipList.Insert(10, 2)

	if node, ok := skipList.Search(2); ok {
		fmt.Printf("Search: index:%v value:%v\n", node.Index(), node.Value())
	}

	if nodes, ok := skipList.SearchRagne(4, 11); ok {
		for _, node := range nodes {
			fmt.Printf("SearchRagne: index:%v value:%v\n", node.Index(), node.Value())
		}
	}

	skipList.ForEach(func(node *SkipNode) bool {
		fmt.Printf("ForEach: index:%v value:%v\n", node.Index(), node.Value())
		return true
	})
}
