package hello_go

import (
    "encoding/json"
    "fmt"
	"log"
	"net/http"
)

type comment struct {
    PostID int `json:"postId"`
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Body string `json:"body"`
}

func callPlaceholderAPI() {
    jsonplaceholderAddress := "http://jsonplaceholder.typicode.com/comments"
	
    var myComments []comment
    
    getJson(jsonplaceholderAddress, &myComments)
    
    fmt.Printf("%v", myComments[0].Email)    
}

func getJson(url string, target interface{}) error {
    res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
    defer res.Body.Close()
    
    return json.NewDecoder(res.Body).Decode(&target)
}
