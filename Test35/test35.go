/*Go no tiene clases. De todas maneras, puedes definir métodos para estructuras.
El receptor del método aparece en su propia lista de parámetros entre la palabra func y el nombre del método.*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

//Result: 5
