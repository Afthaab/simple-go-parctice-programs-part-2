package main

import "fmt"

func main() {
	fmt.Println("Enter the marks : ")
	var marks float32
	fmt.Scanln(&marks)
	if marks > 50 {
		fmt.Println("You have Passed !")

	} else {
		fmt.Println("You have failed")
	}
}
