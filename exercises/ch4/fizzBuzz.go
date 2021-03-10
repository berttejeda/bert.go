package main

import "fmt"

func main() {
	fmt.Println("Using for loop of format i := 1; i <= 10; i++ and Fizz||Buzz as identifiers")
	for i := 1; i <= 100; i++  {		
		switch true {
			case (i % 3 == 0) && (i % 5 == 0): fmt.Println("FizzBuzz")
			case i % 3 == 0: fmt.Println("Fizz")
			case i % 5 == 0: fmt.Println("Buzz")
			default: fmt.Println(i)
		}		
	}
}