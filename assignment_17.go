package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the Choice : ")
	var choice int
	fmt.Println("1. Addition\n2. Substraction\n3. Multiplication\n4. Division")
	fmt.Scan(&choice)
	fmt.Println("Enter the two variable: ")
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	switch choice {
	case 1:
		add(num1, num2)
	case 2:
		mul(num1, num2)
	case 3:
		sub(num1, num2)
	case 4:
		div(num1, num2)
	default:
		fmt.Println("You have Entered wrong choice please try again")
	}
}
func add(a int, b int) {
	var sum int
	sum = a + b
	fmt.Println("The sum of the two numbers: ", sum)

}
func mul(a int, b int) {
	var mul int
	mul = a * b
	fmt.Println("The sum of the two numbers: ", mul)

}
func sub(a int, b int) {
	var sub int
	sub = a - b
	fmt.Println("The sum of the two numbers: ", sub)

}
func div(a int, b int) {
	var div int
	div = a / b
	fmt.Println("The sum of the two numbers: ", div)

}
