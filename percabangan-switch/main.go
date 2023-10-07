package main

import "fmt"

func main() {
	//contoh multicase swicht case statment
	hari := "minggu"

	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println("weekday")
	case "sabtu", "minggu":
		fmt.Println("weekand")

	}

}
