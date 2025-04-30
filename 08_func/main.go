package main

import "fmt"

// outside int -> for return valuer
func add(a int, b int) int {
	return a + b
}

func getLang() (string, string, string) {
	return "golang", "JS", "C"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	sum := add(1, 2)
	fmt.Println(sum)

	fmt.Println(getLang())
	l1, _, _ := getLang()
	fmt.Println(l1)

	// passing fn. inside another fn.
	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	fn2 := processIt2()
	fmt.Println(fn2(6))
}
