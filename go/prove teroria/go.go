package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.Log(-1.0)
	var i int
	for {
		i++
		fmt.Print(i," Ba", n, "a\n")
	}
}
