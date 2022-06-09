package main

import (
	"runtime"
	"time"
)

var _ = runtime.GOMAXPROCS(1)

var a, b int

func u1() {
	a = 1
	b = 2
	println("corriendo u1")
}

func u2() {
	a = 3
	b = 4
	println("corriendo u2")
}

func p() {
	println(a)
	println(b)
}

func main() {
	go u1()
	go u2()
	go p()
	time.Sleep(1 * time.Second)
}
