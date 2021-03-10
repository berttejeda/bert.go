package main

import "fmt"

func main() {
	fmt.Print("Enter degrees in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Print("Degrees in Celsius: ")
	fmt.Println((input - 32) * 5/9)
}