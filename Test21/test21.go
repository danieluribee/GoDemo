/*Un map asocia un valor a cada llave.
Antes de usar un map, hay que crearlo con make (en vez de con new); y el map nil por definición está vacío y no se le puede añadir nada.*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

//Result: {40.68433 -74.39967}
