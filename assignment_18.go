package main

import "fmt"

func main() {
	fmt.Println("Enter the marks for Written Test, Lab Exam, Assignments Respectively")
	var write, lab, assign int
	fmt.Scan(&write, &lab, &assign)
	var grade float32
	grade = float32(write)*70/100 + float32(lab)*20/100 + float32(assign)*10/100
	fmt.Println("The Grade of the Student is : ", grade)

}
