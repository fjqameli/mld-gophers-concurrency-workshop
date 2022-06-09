package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i <= 10000; i++ {

		}
	}()
	fmt.Println("Dropin mic")
	go func() {
		for i := 0; i <= 255; i++ {
			time.Sleep(100 * time.Millisecond)
		}
	}()
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
