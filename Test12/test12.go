//La sentencia if tiene el mismo aspecto que el de C o Java, solo que los paréntesis () no hacen falta (no son ni opcionales), y las llaves {} son obligatorias. (Te suena de algo?)

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

//Result: 1.4142135623730951 2i
