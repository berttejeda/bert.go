package main

import "fmt"

func main() {
	x := "Hello, World"
	fmt.Println(x)
	// When changing the value of an existing variable, you don't need to reinitialize it
	// i.e. don't right x := "Hello again, World", as you'll encounter an error similar to:
	// no new variables on left side of :=
	x = "Hello again, World"
	fmt.Println(x)
}