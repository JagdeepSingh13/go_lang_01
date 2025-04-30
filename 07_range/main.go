package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	fmt.Println(nums)

	sum := 0
	for index, num := range nums {
		fmt.Println(index, num)
		sum = sum + num
	}
	fmt.Println("sum:", sum)

	m := map[string]string{"name": "JSingh", "val": "v1"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// c -> unicode point rune
	// i -> starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c, string(c))
	}
}
