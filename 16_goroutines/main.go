package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	// defer -> runs after fn. complete
	defer w.Done()
	fmt.Println("Doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)

		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}

	// time.Sleep(time.Second * 2)
	wg.Wait()
}
