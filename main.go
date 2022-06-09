package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

type Player struct {
	name  string
	table chan *Ball
}

func main() {
	table := make(chan *Ball)
	p1 := Player{name: "ping", table: table}
	p2 := Player{name: "pong", table: table}
	go p1.Play()
	go p2.Play()
	p1.Serve()

}

func (p *Player) Serve() {
	p.table <- new(Ball)
	time.Sleep(1 * time.Second)
}

func (p *Player) Play() {
	for {
		ball := <-p.table
		ball.hits++
		fmt.Println(p.name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		p.table <- ball
	}
}
