package main

import "fmt"

func main() {
	// declare 1
	var names [3]string
	names[0] = "andi"
	names[1] = "gg"
	names[2] = "gaming"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// declare 2
	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)

	// function len untuk menghitung panjang array
	var pan = [20]string{
		"andi",
	}
	fmt.Println(len(pan)) // result 20
}
