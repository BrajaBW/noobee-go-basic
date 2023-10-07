package main

import (
	"fmt"
)

func main() {

	// operator matematika

	// fmt.Println(1 + 1)
	// fmt.Println(2 - 1)
	// fmt.Println(1 - 9)
	// fmt.Println(6 / 2)
	// fmt.Println(9 % 2)

	// num1 := 10
	// num2 := 2
	// fmt.Println(num1 + num2)
	// fmt.Println(num1 - num2)
	// fmt.Println(num1 - num2)
	// fmt.Println(num1 / num2)
	// fmt.Println(num1 % num2)

	// OPertator INcrement dan descremnet

	i := 5 // 5 +1
	i++

	j := 6 // 6 -1
	j--

	fmt.Println(i)
	fmt.Println(j)
	// operator penugasan

	var a, b, c, d, e = 5, 10, 15, 20, 25

	var x = 1
	a += 1
	b -= 1
	c *= 1
	d /= 1
	e %= 1
	fmt.Println(x)
	fmt.Printf("nilai : %d\n", a)
	fmt.Printf("nilai : %d\n", b)
	fmt.Printf("nilai : %d\n", c)
	fmt.Printf("nilai :%d\n", d)
	fmt.Printf("nilai :%d\n", e)

}
