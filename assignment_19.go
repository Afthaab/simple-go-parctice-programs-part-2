package main

import "fmt"

func main() {
	fmt.Println("Enter the Annumal Income")
	var income int
	fmt.Scan(&income)
	var tax int
	if income <= 250000 && income >= 0 {
		fmt.Println("You Dont need pay tax Enjoy")
	} else if income <= 500000 {
		tax = income * 5 / 100
		fmt.Println("You have to pay : ", tax)
	} else if income <= 1000000 {
		tax = income * 20 / 100
		fmt.Println("You have to pay : ", tax)
	} else if income <= 5000000 {
		tax = income * 30 / 100
		fmt.Println("You have to pay : ", tax)
	} else {
		fmt.Println("your amount is above 50,00,000 ")
	}

}
