package main
import "fmt"
func main()  {
	fmt.Println("Enter the day")
	var day int
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")

	default:
		fmt.Println("In Earth this there are only 7 days goto mars if you need more")
	}

	
}