package main

import "fmt"

func main() {
	fmt.Println(1 + 1)

	var name string = "JSingh"
	var name1 = "John"

	// short-hand syntax
	name2 := "Jai"

	var name3 string
	name3 = "golang"

	const user string = "Zero"

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(host, port)
}
