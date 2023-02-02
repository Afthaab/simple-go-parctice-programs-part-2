package main

import "fmt"

func main() {
	fmt.Println("Enter the 6 subjects marks: ")
	var mark1, mark2, mark3, mark4, mark5, mark6 float32
	fmt.Scan(&mark1, &mark2, &mark3, &mark4, &mark5, &mark6)
	var total = mark1 + mark2 + mark3 + mark4 + mark5 + mark6
	var avg = total / 6
	if avg > 90 && avg <= 100 {
		fmt.Println("You have got A Grade")
	} else if avg > 80 {
		fmt.Println("You have got B Grade")
	} else if avg > 70 {
		fmt.Println("You have got C Grade")
	} else if avg > 60 {
		fmt.Println("You have got D Grade")
	} else if avg > 50 {
		fmt.Println("You have got E Grade")
	} else if avg < 50 {
		fmt.Println("You have Failed Go Study")
	} else {
		fmt.Println("You have entered wrong")
	}
}
