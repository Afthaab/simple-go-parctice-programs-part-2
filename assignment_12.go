package main

import "fmt"

func main() {
	fmt.Println("Enter the limit of the Araay : ")
	var limit int
	fmt.Scan(&limit)
	var array1 [20]int
	fmt.Println("Enter the Array Elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&array1[i])
	}
	// temp := 0
	for i := 0; i < limit; i++ {
		for j := i + 1; j < limit; j++ {
			if array1[j] > array1[i] {
				array1[i],array1[j]=array1[j],array1[i]
			}

		}

	}
	fmt.Println("The Descending Order is : ")
	for i := 0; i < limit; i++ {
		fmt.Println(array1[i])

	}

}
