package main

import "fmt"

func main() {
	fmt.Println("Enter the limit")
	var limit int
	var array [50]int
	sum := 0
	fmt.Scan(&limit)
	fmt.Println("Enter the Array Elements")
	for i := 1; i <= limit; i++ {
		fmt.Scan(&array[i])
		if array[i]%2 != 0 {
			sum = sum + array[i]
		}
	}
	fmt.Println("The sum of the Odd numbers in the array is = ", sum)
}
