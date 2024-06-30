// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	fmt.Println("")
// 	ch := make(chan int)
// 	for j := 0; j < 5; j++ {
// 		wg.Add(2)
// 		go func() {
// 			i := <-ch //receiving the data from the channel
// 			fmt.Println(i)
// 			defer wg.Done()
// 		}()
// 		go func() {
// 			ch <- 42 //giving value to the channel
// 			defer wg.Done()
// 		}()
// 	}

// 	wg.Wait()

// }

//case-2

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	fmt.Println("")
// 	ch := make(chan int)
// 	go func(ch <-chan int) { ////this goroutine is only allowing channel to receive the data
// 		i := <-ch //receiving the data from the channel
// 		fmt.Println(i)
// 		defer wg.Done()
// 	}(ch)
// 	for j := 0; j < 5; j++ {
// 		wg.Add(2)

// 		go func(ch chan<- int) { //this goroutine is only allowing channel to send data
// 			ch <- 42 //giving value to the channel
// 			defer wg.Done()
// 		}(ch)
// 	}

// 	wg.Wait()

// }

//THIS WILL THROW THE ERROR BECAUSE IT ONLY HAS ONE RECEIVER AND MULTIPLE SENDERS
//BY DEFAULT WE ARE WORKING WITH UNBUFFERED CHANNELS WHICH MEANS ONLY ONE MESSAGE
//CAN BE IN THE CHANNEL AT ONE TIME

//------------CASE 3 - TO AVOID ABOVE MENTIONED PROBLEM USE BUFFER CHANNELS------------

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("")
	ch := make(chan int, 50) //buffered channel

	wg.Add(2)
	go func(ch <-chan int) { ////this goroutine is only allowing channel to receive the data
		for i := range ch { //for{
			//  if i,ok:=<-ch;ok{
			// 	fmt.Println(i)
			//  } else{
			// 	break
			//  }
			fmt.Println(i) //}
		}
		defer wg.Done()
	}(ch)
	go func(ch chan<- int) { //this goroutine is only allowing channel to send data
		ch <- 42 //giving value to the channel
		ch <- 27
		close(ch) //closing the channel
		defer wg.Done()
	}(ch)

	wg.Wait()

}
