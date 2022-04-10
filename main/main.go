package main

import (
	"fmt"

	"github.com/AldieNightStar/gox"
)

func main() {
	s := gox.Func("Add", "r *Point", "a, b int", "int",
		"return a + b",
	)
	fmt.Println(s)
}
