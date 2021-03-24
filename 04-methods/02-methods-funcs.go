// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Hypot(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{5, 12}
	fmt.Println(Hypot(v))
}
