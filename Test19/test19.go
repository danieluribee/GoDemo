/*Los literales para estructuras denotan un valor nuevo reservado en memoria con los valores de los campos indicados por orden.
Se puede especificar un subconjunto de los campos usando la notación Campo:. (Y entonces el orden no importa.)
El prefijo especial & obtiene un puntero a la nueva estructura.*/

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}

//Result: {1 2} &{1 2} {1 0} {0 0}
