package main

import (
	"fmt"
)

// func main() {

// 	var fs = [4]func(){}
// 	for i := 0; i < 4; i++ {
// 		defer fmt.Println("defer i = ", i)
// 		defer func() { fmt.Println("defer_closer i = ", i) }()
// 		fs[i] = func() { fmt.Println("closer i = ", i) }
// 	}

// 	for _, f := range fs {
// 		f()
// 	}
// }

func main() {
	a := 10
	b := 20
	X(a, b) // 值传递
	fmt.Println(a, b)

	s := []int{10, 20}
	Y(s) // 引用传递
	fmt.Println(s)

	Z(&a) // 地址传递
	fmt.Println(a)

	f := A // 指向一个函数
	f()    // 调用指向的函数

	ff := func() {
		fmt.Println("匿名函数")
	}
	ff()

	// 闭包函数
	fff := closure(10)
	fmt.Println(fff(1))
	fmt.Println(fff(2))

	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i * 10) //注意: 打印值不变,都是30
		}()
	}

	C()
	panicFunc()
	C()
}

// 函数返回值写法
// 无返回值: func A() {...}
// 单返回值: func A() int {...}
// 多返回值1(完全体): func A() (x int, y int, z int){...}
// 多返回值2(同类型简写): func A() (x,y,z int){...} --> 函数体自带 x,y,z变量, return可简写成return
// 多返回值2(不写名称只写类型): func A() (int, int, int){...} --> 函数体必须指定return变量名(return a,b,c)
// 不定长变参: func A(a ...int){...} --> 变参必须写在最后, 变参是值拷贝, 不会修改原来的值

func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	fmt.Println("Func A")
	return
}

func B() (x int, y int, z int) {
	a, b, c := 1, 2, 3
	return a, b, c
}

// 值拷贝
func X(s ...int) {
	s[0] = 1
	s[1] = 2
	fmt.Println(s)
}

// 引用
func Y(s []int) {
	s[0] = 1
	s[1] = 2
	fmt.Println(s)
}

// 指针参数
func Z(a *int) {
	*a = 2
	fmt.Println(*a)
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func panicFunc() {
	// 注册defer函数
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
