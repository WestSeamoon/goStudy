package main

import (
	"fmt"
)

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

	fmt.Println("请输入需要查找的元素：")
	var target int
	fmt.Scan(&target)

	fmt.Printf("针对 %d 进行二分查找：\n", target)
	idx := binarySearch(list, target)
	if idx == -1 {
		fmt.Printf("数组中没有%d这个元素", target)
	} else {
		fmt.Printf("%d 的索引位置为： [%d]\n", target, idx)
	}
	fmt.Println("----------------------------------------")

	//没有成功获取输入的历程
	//reader := bufio.NewReader(os.Stdin)
	//string, _ := reader.ReadString('\n')
	//name := strings.TrimSpace(string) //去掉字符串两端的空格
	//lists := strings.Split(" ", string)

	//char, err := strconv.Atoi(name)
	//if err != nil {
	//fmt.Println(err)
	//}
	//string1, _ := reader.ReadString(' ')
	//name1 := strings.TrimSpace(string1) //去掉字符串两端的空格

	//char1, _ := strconv.Atoi(name1)

	//fmt.Println(string, name, char, string1, name1, char1)

	//list := make([]int, lens)
	//lists := strings.Split(" ", string)
	//fmt.Println(len(lists), lists)

	//二分查找，查找随机数
	/*
		list := make([]int, 1_000_000)
		for i := 0; i < 1_000_000; i++ {
			list[i] = i
		}
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 20; i++ {
			v := rand.Intn(1_000_000) + 1
			fmt.Printf("针对 %d 进行二分查找：\n", v)
			idx := binarySearch(list, v)
			fmt.Printf("%d 的索引位置为： [%d]\n", v, idx)
			fmt.Println("----------------------------------------")
		}
	*/
}

func binarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	step := 0
	var mid, guess int
	for {
		step = step + 1
		if low <= high {
			mid = (low + high) / 2
			guess = list[mid]
			if guess == target {
				fmt.Printf("共查找了 %d 次\n", step)
				return mid
			}
			if guess < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if guess != target {
				fmt.Printf("共查找了 %d 次\n", step)
				return -1
			}
			return mid
		}
	}

}
