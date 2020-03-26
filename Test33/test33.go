/*Probablemente ya sabías que pinta iba a tener el switch.
Por defecto, al final de un case se sale del switch, a menos que haya una instrucción fallthrough.*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

//Result: Go runs on OS X.
