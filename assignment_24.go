	package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter Your Choice : ")
	fmt.Println("1. String Comparision \n2.Swapping two numbers without using the temperory variable")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		stringcomp()
	case 2:
		swap()

	}

}
func stringcomp() {
	var str1, str2 string
	fmt.Println("Enter the String 1")
	fmt.Scan(&str1)
	fmt.Println("Enter the String 2")
	fmt.Scan(&str2)
	value := strings.Compare(str1, str2)
	if value == 0 {
		fmt.Println("Strings are Equal")
	} else {
		fmt.Println("Strings are not Equal")
	}

}

func swap() {
	var num1, num2 int
	fmt.Println("Enter the two numbers : ")
	fmt.Scan(&num1, &num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("The Swapped Numbers are : ")
	fmt.Println("Number 1: ", num1, " Number 2: ", num2)

}
