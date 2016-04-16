package main

import (
	"fmt"
	"time"
)

type parent struct {
	boy  int
	girl int
}

func (p parent) writeAges() {
	fmt.Println("boy age", p.boy, "girl age", p.girl)
}

type child struct {
	parent
	toys int
}

func exampleUsageOfStructs() {
	myParent := parent{5, 12}

	myParent.writeAges()

	aChild := child{parent{4, 8}, 4}

	aChild.writeAges()

	aChild.parent.writeAges()

	temp := time.Now()

	fmt.Println(temp.String())
}
