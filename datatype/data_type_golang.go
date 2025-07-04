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

	var a float32 = 3.142
	var b float32 = 2.718

	fmt.Printf("Float Value of A: %g\n", a)
	fmt.Printf("Float Value of B: %g\n", b)

	/* -------------------------- Arithmetic Operation -------------------------- */

	fmt.Printf(
		"Addition: %g + %g = %g\n",
		a, b, a+b,
	)

	fmt.Printf(
		"Subtraction: %g - %g = %g\n",
		a, b, a-b,
	)

	fmt.Printf(
		"Multiplication: %g * %g = %g\n",
		a, b, a*b,
	)

	fmt.Printf(
		"Division: %g / %g = %g\n",
		a, b, a/b,
	)

	/* -------------------------------- Booleans -------------------------------- */

	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	str3 := "GeeksforGeeks"

	result1 := str1 == str2
	fmt.Println(result1)

	result2 := str1 == str3
	fmt.Println(result2)

	/* --------------------------------- String --------------------------------- */

	// str variable which stores strings
	str := "GeeksforGeeks"
	var str4 string = "STRING_"
	var str5 string = "Concatenation"

	// Display the length of the string
	fmt.Printf("Length of the string is:%d",
		len(str))

	// Display the string
	fmt.Printf("\nString is: %s", str)

	// Display the type of str variable
	fmt.Printf("\nType of str is: %T", str)

	fmt.Println("\nNew string :", str4+str5)

}
