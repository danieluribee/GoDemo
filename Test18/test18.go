/*Go tiene punteros, pero no tiene aritmética de punteros.
Si tienes un puntero a una estructura, los campos se acceden igual, la indirección a través del puntero es transparente.*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

//esult: {1000000000 2}
