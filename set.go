package main


type set struct {
    data []int
    count int
}

func NewSet() *set {
    s := new(set)
    s.count = 0
    s.data = make([]int,10)
    return s
}

func (s *set) add(item int) {
    s.data[s.count] = item
    s.count++
}

func (s *set) size() int{
    return s.count
}

func (s *set) contains(item int) bool{
    for i := 0; i < s.count; i++ {
        if s.data[i] == item {
            return true
        }
    }
    return false
}

func (s *set) Iter() <-chan int {
    ch := make(chan int);
    go func () {
        for i := 0; i < s.count; i++ {
            ch <- s.data[i]
        }
        // current := b.First
        // fmt.Println( current)
        // for current.next != nil {
        //     fmt.Println( current)
        //     fmt.Println( current.item)
        //     fmt.Println( current.next)
        //     ch <- current.item
        //     current = current.next
        // }
        
        // ch <- current.item
        // current = current.next
        // ch <- current.item
        close(ch)
    } ();
    // fmt.Println( ch)
    return ch
}


