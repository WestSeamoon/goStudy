package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

func worker() {
	n := 0
	next := time.After(time.Second) //定时channel，定时了1秒，该通道在指定时间后会接收到一个值，发送该值的goroutine是go运行时的一部分
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}

func workerupdate() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is", pos)
			next = time.After(time.Second)
		}
	}
}

type command int

const (
	right = command(0)
	left  = command(1)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond //time.Millisecond用于以整数毫秒计数的形式查找持续时间
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func main() {
	//go worker() //这个函数作用相当于每秒打印一个值，值还是自增的
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}

/*
//var mu sync.Mutex

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) Visitlink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	//mu.Lock()
	//defer mu.Unlock() //解锁通常用这种形式

}
*/
