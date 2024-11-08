package hashtab

import (
	"fmt"
	"hash/fnv"
	"strings"
)

type Item struct {
	key   string
	value string
	next  *Item
}

type HashTab struct {
	sizeArr int
	tabl    []*Item
}

func NewHashTab(size int) *HashTab {
	return &HashTab{
		sizeArr: size,
		tabl:    make([]*Item, size),
	}
}

func (ht *HashTab) Hash(itemKey string) int {
	h := fnv.New32a()
	h.Write([]byte(itemKey))
	return int(h.Sum32()) % ht.sizeArr
}

func (ht *HashTab) AddHash(key, value string) {
	index := ht.Hash(key)
	if ht.tabl[index] != nil && ht.tabl[index].key == key {
		fmt.Printf(" ключ '%s' уже существует. значение не добавлено.\n", ht.tabl[index].key)
		return
	}
	newNode := &Item{key: key, value: value, next: ht.tabl[index]}
	ht.tabl[index] = newNode
}

func (ht *HashTab) KeyItem(key string) {
	index := ht.Hash(key)
	if ht.tabl[index] != nil && ht.tabl[index].key == key {
		fmt.Printf("ключ: %s значение: %s\n", key, ht.tabl[index].value)
		return
	}
	fmt.Println("такого ключа нет.")
}

func (ht *HashTab) DelValue(key string) {
	index := ht.Hash(key)
	if ht.tabl[index] != nil && ht.tabl[index].key == key {
		ht.tabl[index] = nil
		return
	}
	fmt.Println("такого ключа нет.")
}

func (ht *HashTab) Print() {
	for i := 0; i < ht.sizeArr; i++ {
		if ht.tabl[i] != nil {
			fmt.Printf("ключ: %s занчение: %s\n", ht.tabl[i].key, ht.tabl[i].value)
		}
	}
}

func WorkHash(comand string, ht *HashTab) {
	parts := strings.Fields(comand)
	if len(parts) < 2 {
		return
	}
	switch parts[1] {
	case "HPUSH":
		if len(parts) < 4 {
			return
		}
		ht.AddHash(parts[2], parts[3])
	case "HRETURNVAL":
		if len(parts) < 3 {
			return
		}
		ht.KeyItem(parts[2])
	case "HPOP":
		if len(parts) < 3 {
			return
		}
		ht.DelValue(parts[2])
	case "PRINT":
		ht.Print()
	}
}
