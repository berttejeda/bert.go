package main

import "fmt"

func main() {
	fmt.Println("Using for loop of format i := 1; i <= 10; i++ with Switch Statement")
	for i := 1; i <= 10; i++  {	
		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unsupported Number")
		}
	}
}