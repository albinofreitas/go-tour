// +build OMIT

package main

import "fmt"

// x, y int
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
