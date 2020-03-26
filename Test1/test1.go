/*Form of a function
func function_name( [Parameter list] ) [return_types]{
	body
}*/

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

//Result: 55
