package array

import (
	"fmt"
	"strconv"
	"strings"
)

type Arr struct {
	number   []string
	sizes    int
	capacity int
}

func NewArr(cap int) *Arr {
	return &Arr{
		number:   make([]string, cap),
		sizes:    0,
		capacity: cap,
	}
}

func (a *Arr) reSize() {
	a.capacity *= 2
	newNumber := make([]string, a.capacity)
	copy(newNumber, a.number)
	a.number = newNumber
}

func (a *Arr) AddArr(value string) {
	if a.sizes >= a.capacity {
		a.reSize()
	}
	a.number[a.sizes] = value
	a.sizes++
}

func (a *Arr) AddAtIndex(index int, value string) {
	if index < 0 || index > a.sizes {
		fmt.Println("индекс вне диапазона!")
		return
	}
	if a.sizes >= a.capacity {
		a.reSize()
	}
	for i := a.sizes; i > index; i-- {
		a.number[i] = a.number[i-1]
	}
	a.number[index] = value
	a.sizes++
}

func (a *Arr) returnItemByIndex(index int) {
	if index < 0 || index >= a.sizes {
		fmt.Println("ииндекс вне диапазона!")
		return
	}
	fmt.Println(a.number[index])
}

func (a *Arr) removeItemByIndex(index int) {
	if index < 0 || index >= a.sizes {
		fmt.Println("индекс вне диапазона!")
		return
	}
	for i := index; i < a.sizes-1; i++ {
		a.number[i] = a.number[i+1]
	}
	a.sizes--
}

func (a *Arr) replaceItemByIndex(index int, value string) {
	if index < 0 || index >= a.sizes {
		fmt.Println("индекс вне диапазона!")
		return
	}
	a.number[index] = value
}

func (a *Arr) Size() int {
	return a.sizes
}

func (a *Arr) Print() {
	for i := 0; i < a.sizes; i++ {
		fmt.Print(a.number[i], " ")
	}
	fmt.Println()
}

func WorkMasiv(comand string, arr *Arr) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "MPUSH":
		if len(parts) < 3 {
			return
		}
		arr.AddArr(parts[2])
	case "MADDINDEX":
		if len(parts) < 4 {
			return
		}
		index, err := strconv.Atoi(parts[3])
		if err != nil {
			return
		}
		arr.AddAtIndex(index, parts[2])
	case "MREMOVE":
		if len(parts) < 3 {
			return
		}
		index, err := strconv.Atoi(parts[2])
		if err != nil {
			return
		}
		arr.removeItemByIndex(index)
	case "MREPLACE":
		if len(parts) < 4 {
			return
		}
		index, err := strconv.Atoi(parts[3])
		if err != nil {
			return
		}
		arr.replaceItemByIndex(index, parts[2])
	case "MRETURN":
		if len(parts) < 3 {
			return
		}
		index, err := strconv.Atoi(parts[2])
		if err != nil {
			return
		}
		arr.returnItemByIndex(index)
	case "MSIZE":
		fmt.Println(arr.Size())
	case "PRINT":
		arr.Print()
	}
}
