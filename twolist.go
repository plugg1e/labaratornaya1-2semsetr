package twolist

import (
	"fmt"
	"strings"
)

type NodeT struct {
	number string
	next   *NodeT
	prev   *NodeT
}

type TwoList struct {
	head *NodeT
	tail *NodeT
}

func AddToEndTwo(lists *TwoList, element string) {
	newNode := &NodeT{number: element, next: nil, prev: lists.tail}
	if lists.head != nil {
		lists.tail.next = newNode
		lists.tail = newNode
	} else {
		lists.head = newNode
		lists.tail = newNode
	}
}

func AddToHeadTwo(lists *TwoList, element string) {
	newNode := &NodeT{number: element, next: lists.head, prev: nil}
	if lists.head != nil {
		lists.head.prev = newNode
		lists.head = newNode
	} else {
		lists.head = newNode
		lists.tail = newNode
	}
}

func DeleteToHeadTwo(lists *TwoList) {
	if lists != nil && lists.head != nil {
		deleteNode := lists.head
		lists.head = deleteNode.next
		if lists.head != nil {
			lists.head.prev = nil
		}
		deleteNode.next = nil
	}
}

func DeleteToEndTwo(lists *TwoList) {
	if lists != nil && lists.tail != nil {
		deleteNode := lists.tail
		lists.tail = deleteNode.prev
		if lists.tail != nil {
			lists.tail.next = nil
		} else {
			lists.head = nil
		}
		deleteNode.prev = nil
	}
}

func DeleteValueTwo(lists *TwoList, element string) {
	if lists.head.number == element && lists.head.next != nil {
		temp := lists.head
		lists.head = lists.head.next
		lists.head.prev = nil
		temp.next = nil
		return
	} else if lists.head.number == element && lists.head.next == nil {
		lists.head = nil
		return
	}
	if lists.tail.number == element {
		temp := lists.tail
		lists.tail = lists.tail.prev
		lists.tail.next = nil
		temp.prev = nil
		return
	}
	temp := lists.head
	for temp != nil {
		if temp.number == element {
			temp.prev.next = temp.next
			temp.next.prev = temp.prev
			temp.next = nil
			temp.prev = nil
			return
		}
		temp = temp.next
	}
}

func QuestTwo(ptr *TwoList, element string) {
	newNode := ptr.head
	for newNode != nil {
		if newNode.number == element {
			fmt.Println("Found")
			return
		}
		newNode = newNode.next
	}
	fmt.Println("ошибка")
}

func PrintNodeTwo(ptr *TwoList) {
	newNode := ptr.head
	if newNode == nil {
		fmt.Println("ошиббка лист пустой")
		return
	}
	fmt.Println("List:")
	for newNode != nil {
		fmt.Println(newNode.number)
		newNode = newNode.next
	}
}

func WorkTwoList(comand string, list *TwoList) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "DPUSHHEAD":
		if len(parts) < 3 {
			return
		}
		AddToHeadTwo(list, parts[2])
	case "DPUSHEND":
		if len(parts) < 3 {
			return
		}
		AddToEndTwo(list, parts[2])
	case "DPOPHEAD":
		DeleteToHeadTwo(list)
	case "DRETURN":
		if len(parts) < 3 {
			return
		}
		QuestTwo(list, parts[2])
	case "DPOPEND":
		DeleteToEndTwo(list)
	case "DPOPV":
		if len(parts) < 3 {
			return
		}
		DeleteValueTwo(list, parts[2])
	case "PRINT":
		PrintNodeTwo(list)
	}
}
