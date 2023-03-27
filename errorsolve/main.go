package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name) //新建一个文件
	if err != nil {
		return err
	}
	defer f.Close() //defer关键字，可以确保deferred的动作可以在函数返回前执行

	_, err = fmt.Fprintln(f, "Errors are values.") //用Fprintln（）向f写入内容
	if err != nil {
		//f.Close() //不管写入是否成功，都需要关闭，否则会产生泄漏
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	//f.Close()
	return err
}

type safeWriter struct {
	w   io.Writer //io.write(...) 解释:将每一个参数写入到文件中(言外之意可以有多个参数),但是参数的类型必须是字符串或者是数字,
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbsjinjie(name string) error {
	f, err := os.Create(name) //新建一个文件
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic")
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic")

	return sw.err
}

func main() {
	files, err := ioutil.ReadDir(".") //读取当前文件夹
	//files, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //Exit()传进一个非0值表示错误
	}
	//fmt.Println(files)
	for _, file := range files {
		fmt.Println(file.Name()) //把目录下所有的文件名打印出来
	}

	err1 := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	err2 := proverbsjinjie("proverbsjinjie.txt")
	if err != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

}
