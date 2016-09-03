package hello_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEmptyBag(t *testing.T) {    
    b := NewBag()
    
	assert.True(t, b.Size() == 0, "test new bag has zero objects inside")
    b.Add(1)
    b.Add(2)
    assert.Equal(t, 2, b.Size(), "test bag has two objects inside")
}



func TestBagIteratesThroughAllItems(t *testing.T) {
     b := NewBag()
     
     type bagObject struct {
         Name string
         inBag bool
     }
     
    objectsToGoInBag :=[]bagObject{bagObject{"will", false, }, bagObject{"james", false}, bagObject{"mike", false}}
    
    for _, bagObject := range(objectsToGoInBag) {
        b.Add(bagObject)
    }
    
    count := 0
    for item := range b.Iter() {
        t.Log("range")
        t.Log(item)
        count++
    }
    
      assert.Equal(t, 3, count, "test bag has three objects inside and we can iterate over them")
}

func TestBagIterateEmpty(t *testing.T) {
     b := NewBag()
     
     count := 0
    for item := range b.Iter() {
        t.Log("range")
        t.Log(item)
        count++
    }
    
    assert.Equal(t, 0, count, "test new bag has zero objects inside and we can iterate over them without crash")
}
    
     
func TestBagRemovesItems(t *testing.T) {
    b := NewBag()
    
	assert.True(t, b.Size() == 0, "test new bag has zero objects inside")
    b.Add(1)
    b.Add(2)
    assert.True(t, b.Size() == 2, "test bag with two additions has two objects inside")
    
    b.Remove()
    assert.Equal(t, 1, b.Size(), "test bag with two additions and one removed has one objects inside")
}