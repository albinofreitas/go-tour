// +build OMIT

package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Println("Pi", Pi)

	const GoRules = true
	fmt.Println("Go rules?", GoRules)
}
