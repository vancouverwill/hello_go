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

func readFileIntoGraph(filename string) {
    results, err := ioutil.ReadFile(filename)
   
   if (err != nil) {
       log.Fatal("file loading error filename:" + filename)
   }
   
   stringResults := string(results[:])
   
   re := regexp.MustCompile(`\n`)
   ints := re.Split(stringResults, -1)
   
//    fmt.Println(len(ints))

//    line := 0
//    var g *diGraph
   
//    v := ints[0]
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
       fmt.Printf("%s\n", numbersTrimmed)
    //    re := regexp.MustCompile(`\s*`)
    //    vertices := re.Split(numbersTrimmed, -1)
       
    //    re := regexp.MustCompile(`([\d]{1,2})`)
       re := regexp.MustCompile(`(\b[[:digit:]]{1,2}\b)`)
       
       
    //    vertices := re.FindStringSubmatch(numbersTrimmed)
    vertices := re.FindAllStringSubmatch(numbersTrimmed, 2)

    // fmt.Printf("%v", numbersTrimmed)
       
       if len(vertices) < 2 {
        //    fmt.Println("v:" + vertices[0])
        fmt.Printf("i:%d", i)
        fmt.Printf("vertices %v", vertices)
        //    log.Fatal(vertices)
       }
       fmt.Printf("v:%s  w:%s arrayLength:%d\n\n", vertices[0][0], vertices[1][1], len(vertices))
   }
   
   fmt.Println(g)
   
   
   
   
//    for _, number := range(ints) {
//     //    fmt.Println(string(i) + ":" + number)
//         fmt.Println(number)
//         if (line == 0) {
//             num, err := strconv.Atoi(number)
//             if (err != nil) {
//                 log.Fatal("not a valid number on first line of input txt")
//             }
            
//         }
//         line++
//    }
   
   
}