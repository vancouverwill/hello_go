package hello_go

import (
	"fmt"
	"github.com/vancouverwill/stringutil"
)

func exampleHelloWorld() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("Hello, world.\n")
	fmt.Printf("Hello, world.\n")
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
