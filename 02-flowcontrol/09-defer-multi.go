// +build OMIT

package main

import "fmt"

//last in, first out - LIFO
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
