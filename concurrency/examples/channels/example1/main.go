package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel of type string and specify the buffer size of 2
	// The buffer size of 2 means that you can send data through the channel twice until it blocks.
	c := make(chan string, 2)
	// Create a Goroutine passing the function arguments along with a channel of type string
	go count("sheep", c)

	// If the msg: <- c is still trying to receive a message but the goroutine has already finished then it will result in a deadlock. Try converting the below for loop into an infinite one, see it in action (only works without a `close()`).
	for i := 1; i <= 10; i++ {
		// Create a variable of type string to receive string output from the above created goroutine's output
		// When receiving from a channel, you can create 2 variables, one which will store the data from the channel which must of the same type as the channel
		// Second is a boolean indicating if the channel is open or not.
		msg, open := <-c

		if open {
			fmt.Println(msg)
		} else {
			fmt.Println("Channel is closed: ", i)
		}

	}
	// There exists a much cleaner syntax of this version
	// Here instead of checking for the open and receiving the channel data you can loop over the range of a channel, which will essentially do the same thing.
	// for msg := range c {
	// 	fmt.Println(msg)
	// }
}

func count(animal string, c chan string) {
	for i := 1; i <= 5; i++ {
		// fmt.Printf("%d %s", i, animal)
		c <- fmt.Sprintf("%d %s", i, animal) // Sends the output from the format.Sprintf() command into the channel.
		time.Sleep(time.Millisecond * 500)
	}

	// Never close the channel as a receiver. Here its ok since we are at the sender's end.
	// WHY? If the sender tries to send the data after receiver closes the channel, the program will panic.
	close(c) // Close the channel so the receiving end doesn't get deadlocked when still trying to receive even after the goroutine has ended.
}
