package main

import (
	"fmt"
)

func main() {
	a := "e"
	b := "é"

	fmt.Printf("%v, %v\n", a, b)

	c := []byte(a)
	d := []byte(b)
	fmt.Printf("%v, %v\n", c, d)
	fmt.Printf("%T, %T", c, d)
}
