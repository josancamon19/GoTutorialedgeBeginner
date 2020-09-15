package main

import (
	"fmt"
	"time"
)
// Tickers - These are excellent for repeated tasks
// Timers - These are used for one-off tasks
func tick() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Ticker tick")
	}
}

func main8() {

	go tick()
	// This print statement will be executed before
	// the first `tick` prints in the console, while the bg task gets initiated
	fmt.Println("The rest of my application can continue")

	select {} // avoid our program ends

}
