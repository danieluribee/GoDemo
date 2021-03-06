/*Las constantes numéricas son valores de alta precisión.
Una constante sin tipo adopta el tipo que el contexto requiere.
Intenta imprimir needInt(Big), también.*/

package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*Results:
21
0.2
1.2676506002282295e+29*/
