package main

import "fmt"

func main() {
	var name *string = new(string)
	*name = "Guillermo"
	fmt.Println(name, *name)

	lastName := "Miralles Campillo"
	fmt.Println(&lastName, lastName)

	var pLastName *string = &lastName
	fmt.Println(pLastName, *pLastName)

	lastName = "Villén Campillo"
	fmt.Println(pLastName, *pLastName)

}
