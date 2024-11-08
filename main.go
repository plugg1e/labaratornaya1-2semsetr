package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"project/internal/array"
	"project/internal/hashtab"
	"project/internal/list"
	"project/internal/queue"
	"project/internal/stack"
	"project/internal/tree"
	"project/internal/twolist"
)

func main() {
	fmt.Println("старт программы")

	arr := array.NewArr(5)
	var root *tree.Tree
	myStack := &stack.Stack{}
	newQueue := &queue.Queue{}
	lists := &list.List{}
	listss := &twolist.TwoList{}
	ht := hashtab.NewHashTab(5)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" ")
		scanner.Scan()
		comand := scanner.Text()
		parts := strings.Fields(comand)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "M":
			array.WorkMasiv(comand, arr)
		case "T":
			tree.WorkTree(comand, &root)
		case "S":
			stack.WorkStack(comand, myStack)
		case "Q":
			queue.WorkQueue(comand, newQueue)
		case "L":
			list.WorkList(comand, lists)
		case "D":
			twolist.WorkTwoList(comand, listss)
		case "H":
			hashtab.WorkHash(comand, ht)
		case "EXIT":
			return
		default:
			fmt.Println("оишбка")
		}
	}
}
