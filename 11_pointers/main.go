package main

import "fmt"

func changeNum(num *int) {
	// dereference
	*num = 5
	fmt.Println("in changeNum", *num)
}

func main() {
	num := 1
	changeNum(&num)

	fmt.Println("after changeNum in main", num)
}
