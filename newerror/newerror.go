//简单做个数独游戏
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

//Grid is a Sudoku grid
type Grid [rows][columns]int8

var (
	//ErrBounds ...
	ErrBounds = errors.New("out of bounds")
	//ErrDigit ...
	ErrDigit = errors.New("invalid digit")
)

//SudoKuError
type SudoKuError []error //自定义错误类型名字应以Error结尾，有时候名字就是Error，例如url.Error

//Error returns one or more errors separated by commas.
func (se SudoKuError) Error() string { //符合error这个接口，所以在set（）方法里可以把SudoError类型的errs变量作为error的返回值类型
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

//Set ...
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudoKuError
	if !inBounds(row, column) {
		//return errors.New("out of bounds")
		//改进一下
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row <= 0 || row >= rows {
		return false
	}
	if column <= 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	if digit <= 0 || digit >= 10 {
		return false
	}
	return true
}

func main() {
	var g Grid
	err := g.Set(12, 0, 15)
	if err != nil {
		/*
			switch err {
			case ErrBounds, ErrDigit:
				fmt.Println("Les erreurs de parametres hors limites.")
			default:
				fmt.Println(err)
				//相当于fmt.Println(err.Error())
			}
		*/
		if errs, ok := err.(SudoKuError); ok { //类型断言
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			//for e := range errs {
			//fmt.Printf("%v: %v\n", e, errs[e])
			//}
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}

	//panic保持冷静并继续
	defer func() {
		if ep := recover(); ep != nil { //调用了recover（）函数，导致停止恐慌，程序会继续运行
			fmt.Println(ep)
		}
	}()
	panic("I forgot m towel") //没有因为panic而导致崩溃
}

/*
//简单做个数独游戏
package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

//Grid is a Sudoku grid
type Grid [rows][columns]int8

var (
	//ErrBounds ...
	ErrBounds = errors.New("out of bounds")
	//ErrDigit ...
	ErrDigit = errors.New("invalid digit")
)

//Set ...
func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		//return errors.New("out of bounds")
		//改进一下
		return ErrBounds
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row <= 0 || row >= rows {
		return false
	}
	if column <= 0 || column >= columns {
		return false
	}
	return true
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit: //errors.New使用的是指针实现的，比较的是内存地址，而不是错误包含的文字信息
			fmt.Println("Les erreurs de parametres hors limites.")
		default:
			fmt.Println(err)
		}
		//fmt.Printf("An error occured: %v.\n", err)
		os.Exit(1)
	}
}
*/
