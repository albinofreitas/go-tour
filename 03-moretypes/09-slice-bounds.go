// +build OMIT

package main

import "fmt"

// When slicing, you may omit the high or low bounds to use their defaults instead.
// The default is zero for the low bound and the length of the slice for the high bound.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s)

	fmt.Println(s[1:4])

	fmt.Println(s[:2])

	fmt.Println(s[1:])
}
