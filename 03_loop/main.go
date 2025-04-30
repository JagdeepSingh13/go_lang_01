package main

import (
	"fmt"
	"time"
)

// only for loop in "go"
func main() {
	// while-loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite-loop
	// for {
	// 	fmt.Println(1)
	// }

	// for-loop
	for i := 0; i < 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// range
	for i := range 3 {
		fmt.Println(i)
	}

	// switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its workday")
	}

	// i.(type) -> return the type of "i"
	// type-switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("not int or string", t)
		}
	}
	whoAmI("JSingh")
	whoAmI(69)

}
