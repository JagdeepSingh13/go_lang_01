package main

import "fmt"

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing no.", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func task(done chan bool) {
// 	defer func() { done <- true }()

// 	fmt.Println("processing task...")
// }

// (emailChan <-chan string) to make it only receiving
// (done chan<- bool) to make it send only

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("sending e-mail to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	// messageChan := make(chan string)

	// sending to channel
	// messageChan <- "ping" // blocking

	// receiving from channel
	// msg := <-messageChan

	// fmt.Println(msg)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 2)

	// result := make(chan int)

	// go sum(result, 1, 2)

	// res := <-result // blocking
	// fmt.Println(res)

	// done := make(chan bool)

	// go task(done)

	// <-done // block

	// buffered channel
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// // emailChan <- "j@j.com"
	// // emailChan <- "j@k.com"

	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan)

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// // it means all the e-mails are already in channel
	// fmt.Println("done sending")

	// // !! important !!
	// close(emailChan)

	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "JSingh"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

}
