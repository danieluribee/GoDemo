//Como en C o Java, puedes dejar vacías las sentencias pre y post.

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
