package main

import "fmt"

func main() {
	/* ---------------------------- Integer DataType ---------------------------- */

	// Using 8-bit unsigned int
	var X int = 32767
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int = 23
	fmt.Println(Y+2, Y-2)

	/* -------------------------- Arithmetic Operation -------------------------- */

	fmt.Printf("Addition: %d + %d = %d\n", X, Y, X+Y)
	fmt.Printf("Subtraction: %d - %d = %d\n", X, Y, X-Y)
	fmt.Printf("Multiplication: %d * %d = %d\n", X, Y, X*Y)
	fmt.Printf("Division: %d / %d = %d\n", X, Y, X/Y)
	fmt.Printf("Modulus: %d %% %d = %d\n", X, Y, X%Y)

	/* ----------------------------- Float DataType ----------------------------- */

}
