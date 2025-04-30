package main

import "fmt"

func counter() func() int {
	var count int = 0

	// because count is of outer-scope to this fn.
	// hence value is not deleted after calling it
	return func() int {
		count += 1
		return count
	}
}

func main() {
	inc := counter()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
