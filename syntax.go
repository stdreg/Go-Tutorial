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

	fmt.Println("\nFunction Syntax")
	fmt.Println(greet("World"))
	fmt.Println("Result from 1 + 4 + 9: ", add(1, 4, 9))
	var x, y = swap(1, 11)
	fmt.Println("Swap 1, 11: ", x, y)
	var xx, yy = swap2(2, 22)
	fmt.Println("Swap2 2, 22: ", xx, yy)
	fmt.Println("Variadic-Add: ", variadicAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 22))

	//using function myfilter as type
	filterfunc := func(s string) bool {
		return len(s) == 3
	}
	var input = []string{"yes", "no", "but", "never"}
	fmt.Println("Anonymous Function as filter: ", myfilter(input, filterfunc))

	fmt.Println("The same with closures (2 and 3):")
	filterfunc2 := filterFuncLen(2)
	filterfunc3 := filterFuncLen(3)
	fmt.Println("Anonymous Function as filter: ", myfilter(input, filterfunc2))
	fmt.Println("Anonymous Function as filter: ", myfilter(input, filterfunc3))

	defertest()

}
