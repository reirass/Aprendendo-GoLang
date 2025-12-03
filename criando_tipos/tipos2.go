package main

import "fmt"

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v, %T\n", x, x)
	// conversão de tipos
	b = hotdog(x) // conversão do tipo de x sendo int para tipo hotdog
	fmt.Printf("%v, %T\n", b, b)
}

// fonte para pesquisar mais sobre: golang ref/spec
