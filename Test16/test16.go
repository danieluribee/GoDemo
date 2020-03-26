/*Un struct es una colección de campos.
(Y una declaración que empieza con type declara un tipo.)*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

//Result: {1 2}
