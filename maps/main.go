package main

import "fmt"

func main() {
	myMap := map[string]int{"foo": 27}
	fmt.Println(&myMap, myMap)
	fmt.Println(myMap["foo"])

	myMap["bar"] = -27
	fmt.Println(myMap)

	delete(myMap, "foo")
	fmt.Println(myMap)
}
