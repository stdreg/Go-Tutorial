package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Basic Syntax")
	var number int //decalaration must be typed, value is 0
	fmt.Println("number:", number)
	var number2 = 2 //assigned - no need to write down type
	fmt.Println("number2:", number2)

	var firstname string = "James"
	var name = "Kirk"
	fmt.Println("Fullname:", firstname, name)

	var pointer *string //value is nil
	fmt.Println("pointer (*):", pointer)
	var list []string //value is []
	fmt.Println("list[]:", list)

	var ( // variables can be grouped
		home   = os.Getenv("HOME")
		usr    = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println("Home:" + home)
	fmt.Println("User:" + usr)
	fmt.Println("GoPath:" + gopath)
	test := "Halloa"      //short notation
	const test2 = "World" //constants are untyped by default
	fmt.Println(test, test2)

	const (
		monday = iota //auto-counter - starts for each group of const at 0
		tuesday
		_ //dont use 2
		thursday
	)
	fmt.Println("monday", monday)
	fmt.Println("tuesday", tuesday)
	fmt.Println("thursday", thursday)

	pointerSyntax() //pointer.go also in package main

	fmt.Println("\nStruct Syntax")
	var meter Meter = 2
	fmt.Println(meter, "in cm", MeterToCm(meter))

	myadress := adress{
		street: "Mainstreet",
		city:   "Metropol",
	}
	fmt.Println(myadress)
	secondadress := adress{"Elmstreet", "NewYork"}
	fmt.Println(secondadress)
	admin := user{
		name: "Admin",
		adress: adress{
			street: "Street1",
			city:   "City1",
		},
	}
	fmt.Println(admin)
	thirdadress := PublicAdress{
		Street: "Street2",
		City:   "London",
	}
	fmt.Println(thirdadress)

}
