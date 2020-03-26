/*De hecho, puedes definir un método para cualquier tipo que definas en un paquete, no sólo estructuras.
Aunque no podrás definir un método para un tipo de otro paquete, o para un tipo básico.*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

//Result: 1.4142135623730951
