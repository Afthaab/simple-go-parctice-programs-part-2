package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the String : ")
	var str string
	flag := true
	fmt.Scan(&str)
	length := len(str)
	for i := 0; i < length; i++ {
		j := length - 1 - i
		if str[i] != str[j] {
			flag = false
			break
		}

	}
	if flag == false {
		fmt.Println("The String is not Pallindrome ! ")
	} else {
		fmt.Println("The String is Pallindrome ! ")

	}

}
