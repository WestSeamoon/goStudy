package main

import (
	"fmt"
	"net"
	"sort"
)

//池并发
func main() {
	ports := make(chan int, 100) //设置100个缓冲区，如果没有100，则通道容量只有1
	results := make(chan []int)  //不需要缓冲，results是main函数在用，main只有一个goroutine
	//var wg sync.WaitGroup
	var openports []int
	var closeports []int

	for i := 0; i < cap(ports); i++ { //cap()获取通道容量
		go worker(ports, results)
	}

	//分配工作,如果直接执行，会堵在后面的输出，所以分配工作也要并行
	go func() {
		for i := 1; i < 1024; i++ {
			//wg.Add(1)
			ports <- i
		}
	}()

	//工作结果
	for i := 1; i <= 1024; i++ {
		port := <-results //由于results执行1024次后会停止，所以之前所有关于WaitGroup的内容都不需要了
		if port[0] == 0 {
			closeports = append(closeports, port[1])
		} else {
			openports = append(openports, port[1])
		}
	}
	//wg.Wait()
	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closeports)

	for _, port := range closeports {
		fmt.Printf("%d closed\n", port)
	}

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports chan int, results chan []int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.31.234:%d", p)
		conn, err := net.Dial("tcp", address) //通道返回用ok，如果返回不成功，ok为false
		var result []int
		if err != nil {
			//fmt.Printf("%d Error!!!\n", p)
			result = append(result, 0, p)
			results <- result
			continue //如果有error，程序永远不会停止，因为没有wg.done（），wg永远不会为0
		}
		conn.Close()
		//fmt.Printf("%d 打开了\n", p)
		result = append(result, 1, p)
		results <- result
		//wg.Done()
	}
}

/*
//并发
func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 120; i++ {
		wg.Add(1) //添加或减少等待goroutine的数量
		go func(j int) {
			defer wg.Done()                                //wg.Done()相当于执行wg.Add（-1）
			address := fmt.Sprintf("192.168.31.223:%d", j) //字符串集合使用Sprintf（）234，223
			conn, err := net.Dial("tcp", address)          //连接
			if err != nil {
				fmt.Printf("%s 关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了\n", address)
		}(i)
	}
	wg.Wait() //执行阻塞，直到所有WaitGroup变为0
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n%d seconds\n", elapsed)
}


//非并发
func main() {
	start := time.Now()
	for i := 1; i < 120; i++ {
		address := fmt.Sprintf("192.168.31.249:%d", i) //字符串集合使用Sprintf（）
		conn, err := net.Dial("tcp", address)          //连接
		if err != nil {
			fmt.Printf("%s 关闭了\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了\n", address)
	}
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n%d seconds\n", elapsed)
}
*/
