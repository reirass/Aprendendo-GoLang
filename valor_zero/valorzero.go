package main

import "fmt"

// valores zero:
// - int: 0 | float: 0.0 | bool: false | string: "" |
// pointers, functions, interfaces, slices, channels, maps: nil
// usar := sempre que possível
// usar var para package level scope

// declaração e inicialização
var a int
var b float32
var c string
var d bool
var x int //-> declaração
func main() {
	// x = 10 // inicialização & atribuição
	// fmt.Printf("%v, %T\n", x, x)
	// x = 20 // atribuição
	// fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
