package main
import "fmt"
func main()  {
	fmt.Println("Enter the Limit")
	var limit int
	fmt.Scan(&limit)
	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j," ")
		}
		fmt.Println()		
	}
	
}