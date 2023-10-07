package main

import "fmt"

func main() {
	//contoh operator logika AND. OR ,NOT
	gender := "male"
	age := 17
	var canWork bool

	if gender == "male" && age >= 18 {
		canWork = true
		fmt.Println("selamat anda bisa bekerja")
	} else {
		canWork = false
		fmt.Println("Maaf Persyaratan Kurang")
	}

	fmt.Println(canWork)

	if gender == "male" || age >= 18 {
		canWork = true
		fmt.Println("selamat anda bisa bekerja")
	} else {
		canWork = false

	}

	fmt.Println(!(canWork))
}
