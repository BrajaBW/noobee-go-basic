package main

import "fmt"

func main() {

	//contoh string
	var txt string = "BRAJA"
	fmt.Println(txt)
	//contoh printf
	fmt.Printf("valuenya : %v,tipedata : %T\n", txt, txt)
	// contoh backtick
	var txt2 string = `
	bogor 
	adalah kota hujan hampir
	setiap hari hujan`
	fmt.Println(txt2)
	//integer
	var nilai1 int16 = -125
	var nilai2 uint8 = 33
	var nilai3 uint64 = 125000000
	fmt.Println(nilai1)
	fmt.Printf("%d\n", nilai2)
	fmt.Printf("%d\n", nilai3)
	//float
	var nilai4 float32 = 3.14
	var nilai5 float64 = 3.0000000000000000000000000000000000000000000005
	fmt.Printf("%.1f\n", nilai4)
	fmt.Printf("%f\n", nilai5)
	//bolean
	var brajaganteng bool = true
	var brajajelek bool = false
	fmt.Println(brajaganteng)
	fmt.Println(brajajelek)

}
