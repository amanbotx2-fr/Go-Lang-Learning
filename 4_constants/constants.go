package main

import "fmt"

func main(){
	const name string = "Chandu"

	fmt.Println(name)


	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}