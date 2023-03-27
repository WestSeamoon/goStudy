package main

import (
	"fmt"
	"time"
)

type person struct {
	name, superpower string
	age              int
}

func birthday(p person) person { //函数struct传入类型需明确
	p.age++
	return p
}

func (p *person) birthday() { //函数struct传入类型需明确
	p.age++
}

func main() {
	canada := "Canada"

	var home *string //声明一个指向string类型的指针类型，类型名home
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(*home)

	timy := &person{
		name: "Timothy",
		age:  10,
	}

	(*timy).superpower = "flying"
	timy.superpower = "flying" //可以这么简写，两效果一样

	fmt.Printf("%+v\n", timy)

	t := person{
		name: "Timothy",
		age:  10,
	}
	fmt.Printf("%v\n", birthday(t))
	fmt.Println(t)

	t.birthday()    //点表记法进行调用时知道使用&取得变量的内存地址
	(&t).birthday() //两句话等效
	fmt.Printf("%v\n", t)

	const layout = "Mon, Jan 2, 2006"
	day := time.Now() //这个now方法并没有修改iatime的值，而是返回了一个新的时间
	tomorrow := day.Add(24 * time.Hour)

	fmt.Println(day.Format(layout)) //format用于格式化方法，即用来控制字符串和变量的显示效果
	fmt.Println(tomorrow.Format(layout))

}
