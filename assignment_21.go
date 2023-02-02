package main

import "fmt"

func main() {
	fmt.Println("Enter the Array Limit : ")
	var limit int
	fmt.Scan(&limit)
	fmt.Println("Enter the Array Elements: ")
	var array1 [10]int
	for i := 0; i < limit; i++ {
		fmt.Scan(&array1[i])
	}
	var array2 [10]int
	fmt.Println("The resultant Array")
	for i := 0; i < limit-1; i++ {
		array2[i] = array1[i] * array1[i+1]
		fmt.Println(array2[i])
	}
	a:=10
	fmt.Printf("%T",a)

}
