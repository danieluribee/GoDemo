/*Insertar o actualizar un elemento en un map m:
m[key] = elem
Obtener un elemento:
elem = m[key]
Borrar un elemento:
delete(m, key)
Determinar si una llave está presente con una asignación múltiple:
elem, ok = m[key]
Si la llave key está en m, ok será true, Si no, ok será false y elem será el valor cero del tipo de los elementos.
Análogamente, al obtener un elemento, si la llave no está presente el resultado es el valor cero del tipo de los elementos.*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

/*Results:
The value: 42
The value: 48
The value: 0
The value: 0 Present? false*/
