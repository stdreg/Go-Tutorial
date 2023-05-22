package main

import (
	"fmt"
	"os"
)

/*
var number int       //decalaration must be typed, value is 0
var firstname string //value is ""
var number2 = 2      //assigned - no need to write down type
var name = "James Kirk"
var User *user     //value is nil
var liste []string //value is nill
*/

func main() {
	var ( // variables can be grouped
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println("Home:" + home)
	fmt.Println("User:" + user)
	fmt.Println("GoPath:" + gopath)
	test := "Halloa" //short notation
	const test2 = "World"
	fmt.Println(test + " " + test2)
}
