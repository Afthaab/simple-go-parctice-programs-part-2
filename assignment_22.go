package main

import "fmt"

func main() {
	a, b, n := getarray()
	sum := addarray(a, b, n)
	fmt.Println("The resultant Array : ")
	displayarray(sum, n)
}
func getarray() ([10]int, [10]int, int) {
	var array1 [10]int
	var array2 [10]int
	var limit int
	fmt.Println("Enter the Array Limit : ")
	fmt.Scan(&limit)
	fmt.Println("Enter the Array 1 Elements : ")
	for i := 0; i < limit; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Println("Enter the Array 2 Elements : ")

	for i := 0; i < limit; i++ {
		fmt.Scan(&array2[i])
	}
	return array1, array2, limit
}
func addarray(a [10]int, b [10]int, n int) [10]int {
	var sum [10]int
	for i := 0; i < n; i++ {
		sum[i] = a[i] + b[i]
	}
	return sum

}
func displayarray(sum [10]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(sum[i])

	}

}
