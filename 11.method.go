package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int // 小写,私有属性, 在package内部任何地方都可以访问(非类型方法也可以访问),其它地方不行
}

type B struct {
	Name string
}

type TZ int

func main() {
	a := A{Name: "aaa", age: 10}
	a.Print()
	fmt.Println(a)

	b := B{}
	b.Print()
	fmt.Println(b)

	var tz TZ
	tz.Print()
	(*TZ).Print(&tz) //第二种调用方法

	// var i int
	// i.Print() // 即使i是int类型(TZ是int类型的别名), 也不能调用 TZ类型的Print()方法
}

// 传递值
func (a A) Print() {
	a.Name = "AA"
	fmt.Println(a)
}

// func (a A) Print(int x) { // error: 与 func (a A) Print() {} 函数冲突
// a.Name = "AA"
// fmt.Println(a)
// }

// 传递指针
func (b *B) Print() {
	b.Name = "BB"
	fmt.Println(b)
}

// 即使TZ是int类型, 也可以有专有的操作方法
func (a *TZ) Print() {
	fmt.Println("TZ")
}
