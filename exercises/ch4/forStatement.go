package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Using for loop of format i <= n ")
	for i <= 10 {
		fmt.Print("Number is: ")
		fmt.Print(i)
		if i % 2 == 0 {
			fmt.Println(" (even)")
		} else {
			fmt.Println(" (odd)")
		}
		i++ // can be written as i = i + 1 or i += 1
	}
	fmt.Println("Using for loop of format i := 1; i <= 10; i++ ")
	for i := 1; i <= 10; i++  {
		fmt.Print("Number is: ")
		fmt.Print(i)
		if i % 2 == 0 {
			fmt.Println(" (even, divisible by 2)")
		} else if i % 3 == 0 {
			fmt.Println(" (odd, divisible by 3)")
		} else if i % 4 == 0 {
			fmt.Println(" (even, divisible by 4)")
		} else {
			fmt.Println(" (odd)")
		}
	}	
}