package stack

import (
	"fmt"
	"strings"
)

type NodeS struct {
	person string
	next   *NodeS
}

type Stack struct {
	head *NodeS
}

func push(stack *Stack, element string) {
	newNode := &NodeS{person: element, next: stack.head}
	stack.head = newNode
}

func pop(stack *Stack) {
	if stack.head != nil {
		newHead := stack.head.next
		stack.head.next = nil
		stack.head = newHead
	}
}

func Show(MyList *Stack) {
	newStack := MyList.head
	for newStack != nil {
		fmt.Print(newStack.person, " ")
		newStack = newStack.next
	}
	fmt.Println()
}

func WorkStack(comand string, stack *Stack) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "SPUSH":
		if len(parts) < 3 {
			return
		}
		push(stack, parts[2])
	case "SPOP":
		pop(stack)
	case "PRINT":
		Show(stack)
	}
}
