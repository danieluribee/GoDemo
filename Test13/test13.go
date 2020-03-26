/*Como el for, el if puede empezar con una instrucción corta que se ejecuta antes de la condición.
El alcance de las variables que se declaran en esa instrucción se limita al código del if.*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

//Result: 9 20
