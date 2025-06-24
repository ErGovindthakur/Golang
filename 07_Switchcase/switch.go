package main

import (
	"fmt"
	"time"
)

func main(){

	var num int 

	fmt.Print("Enter the num -: ")

	fmt.Scan(&num)

	// First variant of switch in go
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other -: ",num)
	}

	// second variant of switch in go (multiple condition)
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's working day")
	}

	// third variant of switch

	WhoAmI := func (i interface{})  {
		switch i.(type){
		case int:
			fmt.Println("Integer")
		case float32:
			fmt.Println("Float")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other")
		}
	}
	WhoAmI(12)
}