package main

import "fmt"

func main() {
	var limit int
	var array1 [20][20]int
	var array2 [20][20]int
	var sum [20][20]int
	fmt.Println("Enter the Limit of the Arrry : ")
	fmt.Scan(&limit)
	fmt.Println("Enter the Array1 Elements : ")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scan(&array1[i][j])
		}
	}
	fmt.Println("Enter the Array2 Elements : ")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scan(&array2[i][j])
		}
	}
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			sum[i][j] = array1[i][j] + array2[i][j]
		}

	}
	fmt.Println("The Sum of the two Arrays : ")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Print(sum[i][j]," ")
		}
		fmt.Println()

	}

}
