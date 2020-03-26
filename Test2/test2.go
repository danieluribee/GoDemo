//Una función puede devolver más de un resultado.
//Esta función devuelve dos strings, o cadenas de caracteres.

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

//Result: world hello
