package main

import "fmt"

// package level scope (ou global scope) não pode usar gopher
// para declarar uma variável PLS é necessário usar var
func main() {
	// := operador de DECLARAÇÃO
	// = operador de ATRIBUIÇÃO
	x := 10 // o operador := (gopher) atribui um tipo com base no valor da variável
	z := 10.0
	y := "olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x) // printf permite formação -> %v mostra o valor e %T o tipo da variável
	fmt.Printf("z: %v, %T\n", z, z) // printf permite formação -> %v mostra o valor e %T o tipo da variável
	fmt.Printf("y: %v, %T\n\n", y, y)

	// x = 20 // isso atribui um valor novo para x
	x, h := 40, 30 // isso atribui o valor novo para X e declara uma nova variável 'h' de valor 30
	// ou seja, o operador declaração precisa PELO MENOS declarar uma nova variável, mesmo que ATRIBUA
	// valor a outra já existente
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("h: %v, %T\n", h, h)
}
