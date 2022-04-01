package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	for i := 1; i < 11; i++ {
		defer spew.Dump(i)
	}
	fmt.Println("bag")

}
