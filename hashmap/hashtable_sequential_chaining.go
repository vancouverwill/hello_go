package hashtable

import "fmt"

type sequentialSearchhHashTb struct {
	hashTable
	content []queue
}

type queue struct {
	start *node
	last  *node
	N     int
}

type node struct {
	value int
	next  *node
}

func NewsequentialSearchhHashTb(size int) *sequentialSearchhHashTb {
	h := new(sequentialSearchhHashTb)
	h.size = size
	h.N = 0
	h.content = make([]queue, size)
	return h
}

func (h *sequentialSearchhHashTb) insert(i int) {
	k := h.hash(i)
	fmt.Printf("hash is %d", k)
	if h.content[k].N == 0 {
		// fmt.Println("new insert")
		h.content[k].start = new(node)
		h.content[k].start.value = i
		h.content[k].last = h.content[k].start
	} else {
		oldLast := h.content[k].last.next
		fmt.Println("another insert")
		h.content[k].last = new(node)
		h.content[k].last.value = i
		oldLast.next = h.content[k].last
	}
	h.content[k].N++
	h.N++
}

func (h *sequentialSearchhHashTb) contains(i int) bool {
	k := h.hash(i)
	fmt.Printf("contains hash is %d", k)
	current := h.content[k].start
	for {
		fmt.Println("current", current.value)
		if current.value == i {
			return true
		}
		if current.next == nil {
			break
		}
		current = current.next
	}
	return false
}
