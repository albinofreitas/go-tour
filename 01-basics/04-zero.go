// +build OMIT

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var c complex64
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, c)
}

// bool => false

// string => ""

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr => 0

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64 => 0

// complex64 complex128 => 0 Real and 0 Imaginary Part (0+0i)