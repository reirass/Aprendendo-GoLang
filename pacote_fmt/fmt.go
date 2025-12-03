package main

import "fmt"

// interpreted string literal vs raw string literal

func main() {
	x := 2
	y := 3
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y) // aqui, ambos ainda são int
	z := fmt.Sprint(x, y)        // pega o valor das variáveis e junta numa mesma string
	// caso ambas variáveis NÃO sejam strings, é adicionado um espaço, caso contrário, é necessário adicionar
	// o espaço com " " entre as variáveis
	// z = 2, 3 em formato de string
	fmt.Printf("%v, %T", z, z)
}

// Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
