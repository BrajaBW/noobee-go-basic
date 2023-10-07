package main

import "fmt"

func main() {

	// //operator perbandingan
	// x := 10
	// y := 5

	// fmt.Println("x == y : ",x==y)

	// //percabangan if else
	// num := 1

	// if num > 0{
	// 	fmt.Println("ok") // jika syarat terpenuhi
	// }else{
	// 	fmt.Println("selesai") // jika syarat tidak terpenuhi
	// }

	// //percabangan if - else if - else

	// nilai := 79

	// var grade string

	// if nilai > 80{
	// 	grade = "A"
	// 	}else if nilai < 80{
	// 		grade = "C"
	// 	}else{
	// 		grade = "D"
	// 	}

	// 	fmt.Println("grade saya adalah :",grade)

	//percabangan nested if
	gender := "man"
	age := 21

	var canWork bool
	if gender == "man" {
		if age = 18 {
			canWork = true
		} else {
			canWork = false
		}
	} else {
		if age <= 20 {
			canWork = true
		} else {
			canWork = false
		}
	}

		if !canWork {// jika kondisi canwork false 
			fmt.Println("tidak diterima")

		} else {
			fmt.Println("diteriman")
		}

		// swict case 

		


	}



