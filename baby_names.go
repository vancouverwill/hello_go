package main

import(
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
)

func GetNames(filename string) {
   results, err := ioutil.ReadFile(filename)
   
   if (err != nil) {
       log.Fatal(err.Error)
   }
   
   regYear := regexp.MustCompile(`(Popularity in)\s([0-9]{4})(</h)`)
   years := regYear.FindSubmatch(results)
   
   for i, year := range(years) {
    fmt.Printf("%d %s\n", i, year)
   }
   
   fmt.Printf("years size %d", len(years))
   
   regBabies := regexp.MustCompile(`(<tr align="right"><td>)(\d{1,6})(</td><td>)(\w+)(</td><td>)(\w+)(</td>)`)
   babies := regBabies.FindAllSubmatch(results, -1)
   
   for i, baby := range(babies) {
       for j, babyWords := range(baby) {
            fmt.Printf("%d %d %s\n", i, j, babyWords)
       }
   }
   
}

func TestGetNames() {
    GetNames("./baby_names_data/baby1990.html")
}