package main

import "fmt"

// tipos em go são estáticos. Exemplo: se eu declaro variável x com tipo int, ele
// será int até finalizar o programa

// var x = 10
var x int

func main() {
	// x = 20.5 // 20.5 é um float, x só pode ser int
	fmt.Printf("%v, %T", x, x)
}
