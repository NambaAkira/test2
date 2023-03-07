package main

import "fmt"

func sum(a int) int {
	a = a + 2
	return a
}

func main() {

	fmt.Println("hello_world")
	fmt.Println("good_morning")

	a := 5
	fmt.Println("before: a = ", a)
	fmt.Println("after: a = ", sum(a))

}
