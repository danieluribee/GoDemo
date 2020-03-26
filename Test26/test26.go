/*Al obtener una porción de un slice el resultado es otro slice que apunta al mismo array que el original.
El resultado de evaluar la expresión
s[lo:hi]
es un slice con los elementos desde lo hasta hi-1, ambos incluidos. Por tanto
s[lo:lo]
está vacío y
s[lo:lo+1]
tiene un elemento.*/

package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}

/*Results:
p == [2 3 5 7 11 13]
p[1:4] == [3 5 7]
p[:3] == [2 3 5]
p[4:] == [11 13]*/
