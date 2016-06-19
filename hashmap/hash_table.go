package hashtable

type hashTable struct {
	size int
	N    int
}

func (h hashTable) getSize() int {
	return h.size
}

func (h hashTable) getN() int {
	return h.N
}

func (h hashTable) hash(i int) int {
	return i % h.size
}
