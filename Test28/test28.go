/*El valor cero de un slice es nil.
Una porción así tiene longitud y capacidad 0.*/

package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

/*Results:
[] 0 0
nil!*/
