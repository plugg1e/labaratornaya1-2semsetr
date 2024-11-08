package list

import (
	"fmt"
	"strings"
)

type Node struct {
	number string
	next   *Node
}

type List struct {
	head *Node
}

func AddToHead(ptr *List, element string) {
	newNode := &Node{number: element, next: ptr.head}
	ptr.head = newNode
}

func AddToEnd(ptr *List, element string) {
	newNode := &Node{number: element, next: nil}
	if ptr.head == nil {
		ptr.head = newNode
	} else {
		nextNode := ptr.head
		for nextNode.next != nil {
			nextNode = nextNode.next
		}
		nextNode.next = newNode
	}
}

func DeleteToHead(node *List) {
	if node.head != nil {
		deleteNode := node.head
		node.head = deleteNode.next
		deleteNode.next = nil
	}
}

func DeleteToEnd(node *List) {
	if node.head != nil {
		if node.head.next == nil {
			node.head = nil
		} else {
			deleteNode := node.head
			for deleteNode.next.next != nil {
				deleteNode = deleteNode.next
			}
			deleteNode.next = nil
		}
	}
}

func DeleteValue(ptr *List, element string) {
	newNode := ptr.head
	var previous *Node
	for newNode != nil && newNode.number != element {
		previous = newNode
		newNode = newNode.next
	}
	if newNode == nil {
		fmt.Println("ошибка")
		return
	}
	if previous == nil {
		ptr.head = newNode.next
	} else {
		previous.next = newNode.next
	}
	newNode.next = nil
}

func Quest(ptr *List, element string) {
	newNode := ptr.head
	for newNode != nil {
		if newNode.number == element {
			fmt.Println("найдено")
			return
		}
		newNode = newNode.next
	}
	fmt.Println("ошибка")
}

func PrintNode(ptr *List) {
	newNode := ptr.head
	if newNode == nil {
		fmt.Println("ошибка лситт пустой")
		return
	}
	fmt.Println("Лист:")
	for newNode != nil {
		fmt.Println(newNode.number)
		newNode = newNode.next
	}
}

func WorkList(comand string, list *List) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "LPUSHHEAD":
		if len(parts) < 3 {
			return
		}
		AddToHead(list, parts[2])
	case "LPUSHEND":
		if len(parts) < 3 {
			return
		}
		AddToEnd(list, parts[2])
	case "LPOPHEAD":
		DeleteToHead(list)
	case "LRETURN":
		if len(parts) < 3 {
			return
		}
		Quest(list, parts[2])
	case "LPOPEND":
		DeleteToEnd(list)
	case "LPOPV":
		if len(parts) < 3 {
			return
		}
		DeleteValue(list, parts[2])
	case "PRINT":
		PrintNode(list)
	}
}
