package main

import (
	"fmt"
)

// membuat constant di luar fungsi
const LENGTH int = 10

func main() {
	// constanta di dalam fungsi
	const WIDTH = 20
	const PI = 3.14
	fmt.Println(LENGTH)
	fmt.Println(WIDTH)
	fmt.Println(PI)

	//constanta deklarasi multipel
	const (
		A int = 40
		B     = 3.14
		C     = "BRAJA"
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
