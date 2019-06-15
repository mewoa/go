// 给interface赋值时，当方法接收者有指针时，将值赋值给接口将导致编译失败
// 调用方法时，可以通过指针或值调用。指针可以默认改变值
package main

import "fmt"

type Student struct {
	name string
}

type intef interface {
	M1()
	M2()
}

func (s Student) M1() {
	s.name = "nihao"
}

func (s *Student) M2() {
	s.name = "yh"
}

func main() {
	// 当接收者类型时指针时，对接收者赋值会直接改变他
	t1 := Student{"t1"}

	fmt.Println("M1调用前：", t1.name)
	t1.M1()
	fmt.Println("M1调用后：", t1.name)

	fmt.Println("M2调用前：", t1.name)
	t1.M2()
	fmt.Println("M2调用后：", t1.name)

	// 对接口赋值, 由于Student类型实现接口时接收者有指针，编译失败
	// var i intef = t1
	// i.M2()
	// i.M1()

	// 更改为指针

}
