//Los literales de map son como los de las estructuras, pero las llaves son necesarias (en el ejemplo, "Bell Labs" y "Google").

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

//Result: map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
