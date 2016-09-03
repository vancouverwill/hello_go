package hello_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEmptySet(t *testing.T) {
    s := NewSet()
    assert.Equal(t, 0, s.size(), "test new set has zero objects inside")
}

func TestAddItemAndContains(t *testing.T) {
    s := NewSet()
    s.add(1)
     assert.Equal(t, true, s.contains(1), "test new set has zero objects inside")
}

func TestIter(t *testing.T) {
    s := NewSet()
    s.add(1)
    s.add(2)
    s.add(3)
    
    count := 0
    for item := range s.Iter() {
        t.Log(item)
        count++
    }
    assert.Equal(t, 3, count, "test iterate returns 3 objects")
}

func TestEmptyIter(t *testing.T) {
    s := NewSet()
    
    count := 0
    for item := range s.Iter() {
        t.Log(item)
        count++
    }
    assert.Equal(t, 0, count, "test iterate returns 3 objects")
}