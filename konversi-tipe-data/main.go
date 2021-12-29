package main

import "fmt"

func main() {
	// konversi int
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi to string
	var name = "andi"
	var a = name[0] // ketika mengambil salah satu karakter result nya jadi byte
	var aString = string(a)
	fmt.Println(aString)
}
