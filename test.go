package main

import "fmt"

func sum(a int) int {
	a = a + 2
	return a
}

func multi(a int, b int) int {
	return a * b
}

func main() {

	fmt.Println("hello_world")
	fmt.Println("good_evening")

	a := 5
	b := 4
	fmt.Println("before: a = ", a)
	fmt.Println("after: a = ", sum(a))

	fmt.Println("multi: a * b = ", multi(a, b))

}
