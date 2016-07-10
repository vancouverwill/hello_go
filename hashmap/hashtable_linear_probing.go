package hashtable

type linearProbinghHashTb struct {
	hashTable
	content []bucket
}

type bucket struct {
	value int
}

func NewLPHashMap(size int) *linearProbinghHashTb {
	h := new(linearProbinghHashTb)
	h.size = size
	h.N = 0
	h.content = make([]bucket, size)
	return h
}

func (h *linearProbinghHashTb) insertIntoLPHashMap(i int) {
	k := h.hash(i)
	for {
		if h.content[k].value == 0 {
			break
		}
		k++
	}
	h.content[k] = bucket{i}
	h.N++
}
