package main

import "fmt"

func main(){
	fmt.Println("Exploring function")

	var taskList = []string{"Study","Play","Eat","Sleep","Code"}

	showTask(taskList)
}

func showTask(taskList []string){
	for index, tasks := range taskList{
		fmt.Printf("%d. %s\n",index+1,tasks)
	}
}