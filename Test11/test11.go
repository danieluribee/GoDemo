//Llegados a ese punto se pueden obviar los puntos-y-coma: el while de C se deletrea for en Go.

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//Result: 1024
