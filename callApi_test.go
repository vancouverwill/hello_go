package hello_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetJSON(t *testing.T) {
    jsonplaceholderAddress := "http://jsonplaceholder.typicode.com/comments"
    var myComments []comment
    getJson(jsonplaceholderAddress, &myComments)
    
	assert.True(t, len(myComments) > 0, "test api returned more than zero comments")
}