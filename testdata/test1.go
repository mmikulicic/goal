package main

import (
	"fmt"
)

// main func
func main() {
	var x int = 1 + 2
	var y int = x * 13
	y = y + 1
	x, y = y, x

	fmt.Println(x, y)
}
