// +build OMIT

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Printf("%v %T\n", v, v)

	v.X = 40
	fmt.Printf("%v %T\n", v.X, v.X)
}
