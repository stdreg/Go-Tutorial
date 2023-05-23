package main

import "fmt"

func pointerSyntax() {
	fmt.Println("\nPointer Syntax")
	var a int = 111
	var pa *int = &a // assign adress of a to pa
	fmt.Println("pa", pa)
	fmt.Println("*pa", *pa) //de-referncing
}
