package main

import (
		"fmt"
    "time"
)

func main() {
	// There's a function process() that can handle only 1 request per second.
  // We need to be able to serve 5 requests per second.
  // The task is to send 100 requests, and then wait for their execution.
  // After all requests are processed, we should stop program execution.
  // For simplicity we assume a request is an integer number.

  // TODO: initialize channels


  // Sending requests.
  for i:= 0; i <= 100; i++ {
    requests <- i
  }
  // TODO: show that no more requests are coming.

  fmt.Println("All requests sent, waiting for them to be processed")

  // TODO: wait till all requests are processed

  fmt.Println("All done")
}

func process(requests chan int, done chan bool) {
  // TODO: implement me
  // hint: you can use processOneRequest()
}

func processOneRequest(request int) {
  fmt.Printf("%d - started\n", request)
  time.Sleep(time.Second)
  fmt.Printf("%d - done\n", request)
}
