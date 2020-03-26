/*Un slice (porción) es un puntero a una tabla de valores e incluye una longitud.
[]T es un slice con elementos de tipo T.*/

package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",
			i, p[i])
	}
}

/*Results:
p == [2 3 5 7 11 13]
p[0] == 2
p[1] == 3
p[2] == 5
p[3] == 7
p[4] == 11
p[5] == 13*/
