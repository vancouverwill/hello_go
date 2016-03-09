package main

import (
    "fmt"
    "errors"
    "io/ioutil"
    "log"
    "regexp"
    "strconv"
    "strings"
)

type  diGraph struct {
    V int
    E int
    adj []Bag
}

func NewDiGraph(V int) *diGraph {
    dg := new(diGraph)
    dg.V = V
    dg.E = 0
    dg.adj = make([]Bag, V)
    return dg
}

func (dg *diGraph) addEdge(v int, w int) error {
    if v < 0 || v >= dg.V { return errors.New("v is out of range") }
    if w < 0 || w >= dg.V { return errors.New("w is out of range") }
    
    dg.adj[v].Add(w)
    dg.E++
    return nil
}

func (dg *diGraph) Adj(v int) (<-chan interface{}, error) {
    if v < 0 || v >= dg.V { return nil, errors.New("v is out of range") }
    return dg.adj[v].Iter(), nil
}

func (dg *diGraph) toString() {
    s := ""
    s += strconv.Itoa(dg.V) + " vertices, " + strconv.Itoa(dg.E) + " edges \n"
    
    for v := 0; v < dg.V; v++ {
            s += strconv.Itoa(v) + ": "
            iterable, err := dg.Adj(v)
            if err != nil {
                log.Fatal("index out of range")
            }
            for w:= range(iterable) {
                wint, ok := w.(int)
                if ok {
                    s += strconv.Itoa(wint) + " "
                }                
            }
            s += "\n"
        }
    
    fmt.Printf("%s", s)
}

func readFileIntoGraph(filename string) {
    results, err := ioutil.ReadFile(filename)
   
   if (err != nil) {
       log.Fatal("file loading error filename:" + filename)
   }
   
   stringResults := string(results[:])
   
   regNewLine := regexp.MustCompile(`\n`)
   ints := regNewLine.Split(stringResults, -1)
   
   v, err := strconv.Atoi(ints[0])
    if (err != nil) {
        log.Fatal("not a valid number of vertices on first line of input txt")
    }
   
   g := NewDiGraph(v)
   
   e, err := strconv.Atoi(ints[1])
    if (err != nil) {
        log.Fatal("not a valid number of edges on first line of input txt")
    }
   
   for i := 2; i < 2 + e; i++ {
       numbersTrimmed := strings.TrimSpace(ints[i])
       re := regexp.MustCompile(`(\b[\d]{1,2}\b)`)
       
       vertices := re.FindAllStringSubmatch(numbersTrimmed, 2)
              
       if len(vertices) < 2 {
            log.Fatal("there is a problem with the content at index " + string(i))
       }
    //    fmt.Printf("v:%s  w:%s arrayLength:%d\n\n", vertices[0][0], vertices[1][1], len(vertices))
       
       v, err := strconv.Atoi(vertices[0][0])
        if (err != nil) {
            log.Fatal("not a valid point v")
        }
       w, err := strconv.Atoi(vertices[1][1])
        if (err != nil) {
            log.Fatal("not a valid point w")
        }
       err = g.addEdge(v, w)
       if (err != nil) {
            log.Fatal("not a valid point w")
        }
   }
   
   g.toString()   
}