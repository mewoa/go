package main

func test(x int) (func(), func()) {

	// 返回的函数变量形成了闭包，都是对x的引用
	// test结束后x变量不会销毁，所有的闭包都会引用它
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}
