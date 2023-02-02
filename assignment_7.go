package main
import "fmt"
func main()  {
	fmt.Println("Enter the number of multiples you need")
	var num int
	fmt.Scan(&num)
	// var array[50] int
	for i := 1; i <= 10; i++ {
		fmt.Println("5 * ",i," = ",5*i)
	}
	
}