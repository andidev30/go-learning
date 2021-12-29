package main

import "fmt"

func main() {
	// operasi matematika
	var a = 1
	var b = 2
	var c = a + b
	fmt.Println(c)

	// operasi matematika augmented assignments
	var d = 10
	d += 10
	fmt.Println(d)

	// operasi matematika unary operator
	var e = 2
	e++
	fmt.Println(e)
}
