package hashtable

// type linearProinghHashMap struct {
// 	size    int
// 	N       int
// 	content []bucket
// }

// type bucket struct {
// 	value int
// }

// func NewLPHashMap(size int) *linearProinghHashMap {
// 	h := new(linearProinghHashMap)
// 	h.size = size
// 	h.N = 0
// 	h.content = make([]bucket, size)
// 	return h
// }

// func (h *linearProinghHashMap) insertIntoLPHashMap(i int) {
// 	k := h.hash(i)
// 	for {
// 		if h.content[k].value == 0 {
// 			break
// 		}
// 		k++
// 	}
// 	h.content[k] = bucket{i}
// 	h.N++
// }

// func (h linearProinghHashMap) getSize() int {
// 	return h.size
// }

// func (h linearProinghHashMap) getN() int {
// 	return h.N
// }

// func (h linearProinghHashMap) hash(i int) int {
// 	return i % h.size
// }
