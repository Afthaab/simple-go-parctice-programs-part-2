package main

import "fmt"

func main() {
	fmt.Println("Enter the limit of the two arrays: ")
	var size int
	fmt.Scan(&size)
	var array1 [20]int
	var array2 [20]int
	fmt.Println("Enter the values of the Array 1: ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Println("Enter the Values of the Array 2: ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array2[i])
	}

	var array3 [20]int
	for i := 0; i < size; i++ {
		array3[i] = array1[i]
		array1[i] = array2[i]
		array2[i] = array3[i]
	}
	fmt.Println("The Array 1: ")
	for i := 0; i < size; i++ {
		fmt.Println(array1[i])
	}
	fmt.Println("The Array 2: ")
	for i := 0; i < size; i++ {
		fmt.Println(array2[i])
	}

}
