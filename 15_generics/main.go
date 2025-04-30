package main

import "fmt"

// can use any, comparable, interface{}, int | string with T

func printSlice[T int | string, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	// nums := []int{1, 2, 3}
	names := []string{"go", "ts"}
	printSlice(names, "JSingh")

	myStack := stack[string]{
		elements: []string{"go", "ts"},
	}
	fmt.Println(myStack)

}
