package main

import "fmt"

func main() {
	fmt.Print("Enter the Two numbers")
	var num1 int
	var num2 int
	fmt.Scan(&num1, &num2)
	var sum = num1 + num2
	fmt.Println("The sum of two numbers are: ", sum)
}
