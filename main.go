package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(1)
	for i := 0; i < 9; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)
}
