package main

import "fmt"

func main() {
	// declaration tipe data
	var name string

	name = "Andi"
	fmt.Println(name)

	name = "Orang Ganteng"
	fmt.Println(name)

	// declaration with initial value
	var age = 13
	fmt.Println(age)

	// declaration with initial value not use var
	country := "indonesia"
	country = "Amerika"
	fmt.Println(country)

	// declaration multiple variable
	var (
		firstname = "Andi"
		lastname  = "Ganteng"
	)
	fmt.Println(firstname, lastname)

}
