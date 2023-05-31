package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func add(a, b, c int32) int32 {
	return a + b + c
}

func swap(a, b int) (int, int) { //returning multiple values
	return b, a
}

func swap2(a, b int) (x int, y int) { //return values can have names
	x = b
	y = a
	return //no need to write down what gets returned because its already declared
}

func variadicAdd(a ...int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}

type filterFunction func(string) bool

func myfilter(s []string, filter filterFunction) []string {
	var out []string
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}
	return out
}

//closure for filtering with parameter
func filterFuncLen(length int) filterFunction {
	return func(s string) bool {
		return len(s) == length
	}
}

//defer - calls get executed at the end of the method
func defertest() {
	fmt.Println("start defertest")
	defer fmt.Println("defer1")
	fmt.Println("end defertest")
}
