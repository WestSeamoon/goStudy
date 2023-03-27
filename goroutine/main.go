package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	//downstream <- "" //哨兵值，表示内容发完了，使用close关闭
	close(downstream)
}

//把上游传来的值经过筛选后传给下游
func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			//item := <-upstream
			//if item == "" {
			//downstream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

//使用range关键字读取值，直到通道关闭,与filterGopher函数效果一样
func filterGopherRange(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopherRange(c0, c1)
	printGopher(c1)
}

/*
//channel
func main() {
	c := make(chan int)
	//go func() { c <- 2 }()
	//<-c	//如果没有上面的匿名函数，goroutine一直等待发送或接收时，即被阻塞了，由于永远无法发生，这种情况称为死锁
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	//for i := 0; i < 5; i++ {
		//gopherID := <-c
		//fmt.Println("gopher ", gopherID, "has finished sleeping")
	//}
	//使用select处理多个通道
	timeout := time.After(2 * time.Second) //time.After（）返回一个通道，该通道在指定时间后会接收到一个值，发送该值的goroutine是go运行时的一部分
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
*/

/*
//goroutine
func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	//go sleepyGopher()           //分支线路
	time.Sleep(4 * time.Second) //主线路，如果主线路为1秒，则分支线路也会在1秒内结束
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}
*/
