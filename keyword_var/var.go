package main

import "fmt"

var y = 10 // declarando de forma package level scope, o y pode ser usado
// em todos os codeblocks
// fora do codeblock, o gopher (:=) não pode ser usado, sendo obrigado a usar var

func main() {
	z := 20
	// y := 10 // declarando dentro dessa função, só poderá ser usada nesse escopo
	algo(z)
}

func algo(x int) {
	fmt.Println(y)
	fmt.Println(x)
}
