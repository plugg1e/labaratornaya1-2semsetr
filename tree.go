package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func AddTree(root **Tree, value int) {
	if *root == nil {
		*root = &Tree{value: value, left: nil, right: nil}
		return
	}
	if value <= (*root).value {
		AddTree(&(*root).left, value)
	} else {
		AddTree(&(*root).right, value)
	}
}

func removeTree(root *Tree, value int) {
	if root == nil {
		fmt.Println("такого элемента нету")
		return
	}
	if value <= root.value {
		if value == root.value {
			fmt.Println("есть элемент")
			return
		}
		removeTree(root.left, value)
	} else {
		if value == root.value {
			fmt.Println("есть элемент")
			return
		}
		removeTree(root.right, value)
	}
}

func printTree(root *Tree, level int) {
	if root != nil {
		printTree(root.left, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Println(root.value)
		printTree(root.right, level+1)
	}
}

func countNumNodes(tree *Tree) int {
	if tree == nil {
		return 0
	}
	return 1 + countNumNodes(tree.left) + countNumNodes(tree.right)
}

func checkComplete(tree *Tree, index int, numberNodes int) bool {
	if tree == nil {
		return true
	}
	if index >= numberNodes {
		return false
	}
	return checkComplete(tree.left, 2*index+1, numberNodes) && checkComplete(tree.right, 2*index+2, numberNodes)
}

func WorkTree(comand string, root **Tree) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "TPUSH":
		if len(parts) < 3 {
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			return
		}
		AddTree(root, value)
	case "TREMOVE":
		if len(parts) < 3 {
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			return
		}
		removeTree(*root, value)
	case "TCHECK":
		nodeCount := countNumNodes(*root)
		index := 0
		if checkComplete(*root, index, nodeCount) {
			fmt.Println("это бинарное дерево")
		} else {
			fmt.Println("это не бинарное дерево")
		}
	case "PRINT":
		printTree(*root, 0)
	}
}
