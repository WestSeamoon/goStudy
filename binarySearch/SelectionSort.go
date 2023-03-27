package main

import "fmt"

func main() {
	fmt.Println("请输入数组的长度：")
	var lens int
	fmt.Scan(&lens) //reader.Text()
	//fmt.Printf("%d \n", lens)

	fmt.Println("请输入数组的元素,用空格隔开：")

	list := make([]int, lens)
	for i := 0; i < lens; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(list)

	fmt.Printf("选择排序的结果为：%d\n", SelectSort(list))
}

func SmallerIndex(list []int) int {
	Smaller := list[0]
	index := 0
	for i := range list {
		if Smaller > list[i] {
			Smaller = list[i]
			index = i
		}
	}
	return index
}

func SelectSort(list []int) []int {
	result := []int{}
	count := len(list)
	for i := 0; i < count; i++ {
		index := SmallerIndex(list)
		result = append(result, list[index])
		list = append(list[:index], list[index+1:]...)
	}
	return result
}
