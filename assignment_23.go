package main

import "fmt"

func main() {
	a, b, n := getarray()
	sum := addarray(a, b, n)
	displayarray(sum, n)
}
func getarray() ([10][10]int, [10][10]int, int) {
	var array1 [10][10]int
	var array2 [10][10]int
	var limit int
	fmt.Println("Enter the Array Limit: ")
	fmt.Scan(&limit)
	fmt.Println("Enter the Array 1 Elements: ")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scan(&array1[i][j])
		}
	}
	fmt.Println("Enter the Array 2 Elements: ")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scan(&array2[i][j])
		}
	}
	return array1, array2, limit
}
func addarray(a [10][10]int, b [10][10]int, n int) [10][10]int {
	var sum [10][10]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum[i][j] = a[i][j] + b[i][j]
		}

	}
	return sum
}
func displayarray(sum [10][10]int, n int) {
	fmt.Println("The Resultant Array is: ")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(sum[i][j], " ")

		}
		println()

	}

}
