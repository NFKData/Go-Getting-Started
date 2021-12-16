package main

import "fmt"

func main() {
	const a = 3
	fmt.Println(a + 1)
	fmt.Println(a + 3.14)

	const b int = 3
	fmt.Println(b + 1)
	fmt.Println(float32(b) + 3.14)
}
