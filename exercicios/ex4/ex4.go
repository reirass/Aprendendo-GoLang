package main

import "fmt"

type bolsa int

var x bolsa

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Printf("%v, %T", x, x)

}
