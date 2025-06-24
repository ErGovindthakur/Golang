// exploring constant in go

package main

import (
	"fmt"
	// "go/constant"
)


// Note -> You can also assign constant or variables values outside the main function

const name = "Govind Thakur"
var age = 20

// But you can't do the below work
// name := "Raj"

func main(){

	// Note -> Once you assign the values using constant it can't be change and reassign

	const dob = "16 March"
	fmt.Println(dob)

	// always access the values inside the main method either values declare inside or outside

	fmt.Println(name,"is",age)

	// Explore constant grouping

	const (
		serverName = "Golang"
		port = 5001
		serverMethod = "wireless"
		isWireless = true
	)

	fmt.Println("The server", serverName, "is", serverMethod, "running at port", port)
}