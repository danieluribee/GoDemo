/*La instrucción var declara una lista de variables; tal como en los parámetros de una función, el tipo va al final.
Una declaración de variables puede incluir valores iniciales, uno por variable.
Si éstos estan presentes, se puede omitir el tipo y la variable tendrá el tipo del valor inicial utilizado.*/

package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(x, y, z, c, python, java)
}

//Result: 1 2 3 true false no!
