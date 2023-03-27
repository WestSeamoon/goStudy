package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger //几乎任何东西
	Info    *log.Logger //重要信息
	Warning *log.Logger //警告
	Error   *log.Logger //错误
)

func init() { //init()函数会在main函数之前运行
	log.SetPrefix("ABC: ") //prefix前缀，设置用set，按惯例前缀是大写的

	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666) //打开c.log，没有就创建，有就添加，只能为写，权限为0666
	if err != nil {
		log.Fatalln("无法打开错误文件：", err)
	}
	log.SetOutput(file)

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	/*
		//setFlag的取值log.
		const (
			Ldate = 1 << iota
			Ltime
			Lmicroseconds		//毫秒
			Llongfile			//绝对路径文件名和行号
			Lshorfile			//文件名和行号
			LUTC 				//日期时间转为0时区的
			Lmsgprefix			//flag开启Lmsgprefix的时候，prefix在要打印的内容"1234"的前面
			LstdFlags = Ldate | Ltime
		)
	*/

	Trace = log.New(ioutil.Discard, //ioutil.Discard,返回len（），不会干什么，如果把os.Stdout赋给Discard,那么trace log就会写入Stdout，相当于一个开关变量，写入值就会写log，采用默认值就会写丢
		"Trace:",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO:",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(file, os.Stdout), //向两个地方写入，整合成一个参数
		"WARNING:",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR:",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	//log.Println("1234")

	//log.Fatalln("1234")	//相当于将日志打印后，调用os.exit(1)，用于记录一些重大的错误然后退出程序

	//log.Panicln("1234") //打印panic后程序没有停止

	Trace.Println("鸡毛蒜皮的小事")
	Info.Println("一些特别的信息")
	Warning.Println("这是一个警告")
	Error.Println("出现了故障")
}
