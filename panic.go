// 测试panic和defer的行为，defer以队列方式存储，
// 函数遇到panic，会去处理defer，之后向上传递，打印panic的内容
// 最后，打印函数调用栈
// recover只会捕获最后一个panic
package main

import (
	"fmt"
)

func main() {
	defer callFunc()
	panic("A panic!")
	fmt.Println("After recover")
}

func callFunc() {
	if r := recover(); r != nil {
		fmt.Print("panic caugth, do some op...")
	}
}
