package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var s, sep string，可以使用短声明
	//os.Args//是一个[]string
	//s, sep := "", ""

	/*
		for i := 1; i < len(os.Args); i++ { //i为什么从1开始，因为命令本身就是一个args
			s += sep + os.Args[i]
			sep = " "
		}
	*/

	/*
		//改进
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}

		fmt.Println(s)
	*/

	//最简洁的优化
	fmt.Println(strings.Join(os.Args[1:], " ")) //Jion函数可以把集合字符串转化为一个字符串

	//读取控制台参数
	fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)  //用bufio.NewReader读取用户的输入，里面放系统的标准输入os.Stdin
	text, err := reader.ReadString('\n') //读取用户输入，通过‘\n’换行符表示输入结束了
	if err != nil {
		fmt.Println("读取错误")
	}
	fmt.Printf("Your name is :%s", text)
}
