package queue

import (
	"fmt"
	"strings"
)

type NodeQ struct {
	person string
	next   *NodeQ
}

type Queue struct {
	head *NodeQ
	tail *NodeQ
}

func pushQ(queues *Queue, element string) {
	newNode := &NodeQ{person: element, next: nil}
	if queues.head != nil {
		queues.tail.next = newNode
		queues.tail = newNode
	} else {
		queues.tail = newNode
		queues.head = queues.tail
	}
}

func popQ(queues *Queue) {
	if queues.head != nil {
		newHead := queues.head.next
		queues.head.next = nil
		queues.head = newHead
		if queues.head == nil {
			queues.tail = nil
		}
	} else {
		fmt.Println("ошибка лист пуст")
	}
}

func ShowQ(queues *Queue) {
	newQueue := queues.head
	for newQueue != nil {
		fmt.Print(newQueue.person, " ")
		newQueue = newQueue.next
	}
	fmt.Println()
}

func WorkQueue(comand string, queue *Queue) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "QPUSH":
		if len(parts) < 3 {
			return
		}
		pushQ(queue, parts[2])
	case "QPOP":
		popQ(queue)
	case "PRINT":
		ShowQ(queue)
	}
}
