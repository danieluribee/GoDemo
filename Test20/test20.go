/*La expresión new(T) devuelve un puntero a un nuevo valor de tipo T inicializado a 0.
var t *T = new(T)
or
t := new(T)8*/

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}

/*Results:
&{0 0}
&{11 9}*/
