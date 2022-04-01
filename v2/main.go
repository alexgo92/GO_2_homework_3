package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	for i := 1; i < 11; i++ {
		spew.Dump(i)
	}
}
