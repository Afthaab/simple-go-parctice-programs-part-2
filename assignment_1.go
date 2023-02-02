package main

import "fmt"

func main() {
	fmt.Print("Enter the name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("The given name is: ", name)
}
