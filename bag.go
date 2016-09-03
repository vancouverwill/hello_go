package hello_go


type Bag struct {
    N int
    First *node
}

type node struct {
    item interface{}
    next *node
}

func NewBag() *Bag {
    newbag := new(Bag)
    newbag.First = nil
    newbag.N = 0
    return newbag
}

func (b *Bag) Add(item interface{}) {
    oldFirst := b.First
    b.First = new(node)
    b.First.item = item
    b.First.next = oldFirst
    b.N++
}

func (b *Bag) Remove() interface{} {
    oldFirst := b.First
    b.First = oldFirst
    b.N--
    return oldFirst.item
}

func (b *Bag) Size() int {
    return b.N
}

func (b *Bag) IsEmpty() bool {
    return b.First == nil
}

/*
Iter returns items in no particular order
*/
func (b *Bag) Iter() <-chan interface{} {
    ch := make(chan interface{}, b.N);
    go func () {
        if b.First == nil {
            close(ch)
            return
        }
        current := b.First
        for current.next != nil  {
            ch <- current.item
            current = current.next
        }
        
        ch <- current.item
        // current = current.next
        close(ch)
    } ();
    return ch
}

