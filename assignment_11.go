package main

import "fmt"

func main() {
	fmt.Println("Enter the Size of the Array : ")
	var size int
	fmt.Scan(&size)
	fmt.Println("Enter the Array Elements : ")
	var array1 [20]int
	count := 0
	for i := 0; i < size; i++ {
		fmt.Scan(&array1[i])
		if array1[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("The Number of Even Numbers in the Array are : ", count)

}
