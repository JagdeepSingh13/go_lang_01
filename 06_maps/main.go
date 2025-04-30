package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	// [type of key]type of value

	m := make(map[string]string)

	// setting an element
	m["name"] = "JSingh"
	m["area"] = "back"

	fmt.Println(m["name"], m["area"])
	fmt.Println(len(m))

	delete(m, "area")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	m2 := map[string]int{"price": 10, "qty": 3}
	m3 := map[string]int{"price": 10, "phone": 3}
	fmt.Println(m2)

	// _ -> if that value is not being used
	v, ok := m2["price"]
	if ok {
		fmt.Println("all ok")
		fmt.Println(v)
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(maps.Equal(m2, m3))
}
