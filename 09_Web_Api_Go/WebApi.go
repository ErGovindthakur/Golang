// Writing first time web api in go

package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Welcome to My WebApi")


	http.HandleFunc("/",sayHello)

	http.ListenAndServe(":8080",nil)
}

func sayHello(writer http.ResponseWriter, request *http.Request){
	var greet = "Hello , How are you doing ?"
	fmt.Fprintln(writer,greet)
}