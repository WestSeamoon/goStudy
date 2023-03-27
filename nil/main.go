package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil { //没有这一步，会导致恐慌
		return //直接返回或打印
	}
	p.age++
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return len(s[i]) < len(s[j]) }
	}
	sort.Slice(s, less) //排序
}

func main() {
	var nobody *person
	fmt.Println(nobody)

	nobody.birthday()

	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)

	//nil map
	var soup map[string]int
	fmt.Println(soup == nil)

	measurement, ok := soup["onion"]
	if ok { //false,不会运行
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement) //nil，没有遍历，不会运行
	}

	//未被赋值的接口变量接口类型和值都是nil
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	//当接口类型的变量被赋值后，接口就会在内部指向该变量的类型和值
	var p *int
	v = p //v的值依然是nil，但是类型发生了变化
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	//检验接口类型的内部表示
	fmt.Printf("%#v\n", v)

	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e) //符合String（）方法
}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

//nil之外的另一种选择
func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}
