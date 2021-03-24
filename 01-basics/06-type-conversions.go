// +build OMIT

package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42

	var f float32
	f = float32(i)

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", f, f)

	f2 := 32.9
	i2 := int(f2) // descarta a fração
	fmt.Printf("%v %T\n", i2, i2)

	var s string
	s = string(i) // 42 na tabela ascii representa o *
	fmt.Printf("%v %T\n", s, s)

	s1 := strconv.Itoa(i)
	fmt.Printf("%v %T\n", s1, s1)
}
