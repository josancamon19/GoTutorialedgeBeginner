package main

import (
	"fmt"
	"runtime"
)

func main3()  {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}