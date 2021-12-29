package main

import "fmt"

func main() {
	var nilai = 90
	var absensi = 80

	var lulusNilai = nilai >= 80
	var lulusAbsensi = absensi >= 75

	var lulus = lulusAbsensi && lulusNilai
	fmt.Println(lulus)
}
