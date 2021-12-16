package main

import "fmt"

func main() {
	var i int
	i = 7
	fmt.Println(i)

	var pi float32 = 3.14
	fmt.Println(pi)

	a := "Pepe"
	fmt.Println(a)

	b := true
	fmt.Println(b)

	c := complex(1, 2)
	fmt.Println(c)

	real, imaginary := real(c), imag(c)
	fmt.Println(real, imaginary)
}
