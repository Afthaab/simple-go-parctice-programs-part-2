package main

import "fmt"

func main() {
	fmt.Println("Enter the Limit of the Array: ")
	var limit int
	fmt.Scan(&limit)
	k := 1
	for i := 0; i <= limit; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(k," ")
			k++
		}
		println()

	}
}
