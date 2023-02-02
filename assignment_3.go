package main

import "fmt"

func main() {
	fmt.Println("Enter the Principle : ")
	var principle float32
	fmt.Scanln(&principle)
	fmt.Println("Enter the Rate of Intrest : ")
	var rate float32
	fmt.Scanln(&rate)
	fmt.Println("Enter the Time : ")
	var time float32
	fmt.Scanln(&time)
	var Simple float32
	Simple = (rate * time * principle) / 100
	fmt.Println("The Simple Intrest is : ", Simple)
}
