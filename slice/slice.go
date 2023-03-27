//import "fmt"

//golang提供了多重赋值的特性可以轻松实现变量的交换，变量一，变量二 ：= 变量二，变量一
func convert(s []int) []int {

	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
		//      tmp := s[i]
		//      s[i] = s[length-i-1]
		//      s[length-i-1] = tmp
	}
	return s
}