//Si el tipo del nivel superior es simplemente un nombre de tipo, se puede omitir en los elementos del literal.

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

//Result: map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
