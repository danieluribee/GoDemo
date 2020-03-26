/*Los slices se crean con la función make. Ésta reserva memoria para una tabla inicializada a cero y devuelve un slice que apunta a esa tabla:
a := make([]int, 5)  // len(a)=5
Los slices tienen una longitud y una capacidad. La capacidad es el tamaño máximo que puede abarcar el slice en la tabla subyacente.
Para especificar una capacidad, se puede pasar un tercer argumento a make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5
Un slice se puede alargar, con límite la capacidad, obteniendo una porción más larga:

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*Results:
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]*/
