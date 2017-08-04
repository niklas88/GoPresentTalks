package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var notUsed string // forbidden

	var someFloat float32
	someInt := 42 // uses type inference

	someFloat = someInt // forbidden
	fmt.Println("Need to actually use it:", someFloat)
	// END OMIT
}
