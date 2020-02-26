package main

import "fmt"

func main() {
	// var i int
	// i = 42
	// fmt.Println(i)

	// var f float32 = 3.14
	// fmt.Println(f)

	// firstName := "Anthony"
	// fmt.Println(firstName)

	// b := true
	// fmt.Println(b)

	// c := complex(3, 4)
	// fmt.Println(c)

	// r, im := real(c), imag(c)
	// fmt.Println(r, im)

	// var firstName *string = new(string)
	// *firstName = "Anthony"
	// fmt.Println(*firstName)

	firstName := "Anthony"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)
}
