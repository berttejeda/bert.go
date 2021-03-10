package main

import "fmt"

func main() {
	fmt.Print("Enter number in feet: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Print("Feet in meters: ")
	fmt.Println((input/0.3048))
}