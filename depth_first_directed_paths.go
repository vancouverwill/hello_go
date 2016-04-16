package main 

import (
    "fmt"
    "log"
)

type DepthFirstDirectedPaths struct {
    marked []bool
    edgeTo []int
    s      int      // source vertex
}

func NewDepthFirstDirectedPaths(G diGraph, s int) *DepthFirstDirectedPaths {
    df := new(DepthFirstDirectedPaths)
    df.s = s
    df.marked = make([]bool, G.V)
    df.edgeTo = make([]int, G.V)
    df.dfs(G, s)
    return df
}

func (df *DepthFirstDirectedPaths) dfs(G diGraph, v int) {
    df.marked[v] = true
    iterable, err := G.Adj(v)
        if err != nil {
            log.Fatal("index out of range")
        }
    for w := range(iterable) {
        if !df.marked[w.(int)] {
            df.edgeTo[w.(int)]=v
            df.dfs(G, w.(int))
        }
    }
}

func (df *DepthFirstDirectedPaths) hasPathTo(v int) bool {
    return df.marked[v]
}

// return path from s to v; null if no such path
func (df *DepthFirstDirectedPaths)pathTo(v int) <-chan int {
        if !df.hasPathTo(v) { return nil }
        
        ch := make(chan int, len(df.marked));
        go func () {
            for x := v; x != df.s; x = df.edgeTo[x] {
                ch <- x
            }
            ch <- df.s
            
            // if b.First == nil {
            //     close(ch)
            //     return
            // }
            // current := b.First
            // for current.next != nil  {
            //     ch <- current.item
            //     current = current.next
            // }
            
            // ch <- current.item
            // current = current.next
            close(ch)
        } ();
        return ch
        
        // Stack<Integer> path = new Stack<Integer>();
        // for (int x = v; x != s; x = edgeTo[x])
        //     path.push(x);
        // path.push(s);
        // return path;
    }

func displayDepthFirstDirectedPaths(filename string, s int) {
    g := readFileIntoGraph(filename)
    
    dfs := NewDepthFirstDirectedPaths(g, s)
  
    for v := 0; v < g.V; v++ {
            if dfs.hasPathTo(v) {
                fmt.Printf("%d to %d:  ", s, v);
                for x := range dfs.pathTo(v) {
                    if x == s { 
                        fmt.Printf("-%d", x); 
                     }  else {   
                         fmt.Printf("-%d", x);
                    }
                }
                fmt.Println();
            } else {
                fmt.Printf("%d to %d:  not connected\n", s, v);
            }

        }
}