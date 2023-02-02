package main

import "fmt"

func main() {
	fmt.Println("Enter the Number")
	var num int
	fmt.Scan(&num)
	var flag bool
	for i := 2; i < num; i++ {
		flag = false
		if num%2 == 0 && num > 1 {
			flag = true
		}

	}
	if flag == true {
		fmt.Println("The number is not Prime")
	} else {
		fmt.Println("The number is Prime")
	}

}
