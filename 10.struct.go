package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type person2 struct {
	Name    string
	Age     int
	Contact struct { // 匿名结构
		Phone   string
		Address string
	}
}

type person3 struct {
	string
	int
}

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
	Sex  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	// 简单用法
	a := person{
		Name: "joe",
		Age:  19}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	B(&a)
	fmt.Println(a)
	C(&a)
	fmt.Println(a)
	fmt.Println("--------------")

	// 直接使用地址
	b := &person{
		Name: "joe",
		Age:  19}
	b.Name = "john" // 即使b是地址, 也是使用.符号进行操作
	fmt.Println(b)
	B(b)
	fmt.Println(b)
	C(b)
	fmt.Println(b)
	fmt.Println("--------------")

	// 使用匿名结构
	c := &struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(c)
	fmt.Println("--------------")

	// person2中嵌入匿名结构的用法
	p2 := person2{Name: "john", Age: 18}
	p2.Contact.Phone = "18888888888"
	p2.Contact.Address = "My City"
	fmt.Println(p2)

	p3 := person3{"jack", 88}
	fmt.Println(p3)
	fmt.Println("--------------")

	// 嵌入结构
	tea := teacher{
		Name:  "ttt",
		Age:   30,
		human: human{Sex: 1},
	}
	fmt.Println(tea)
	tea.Sex = 100 // 当teacher有Sex属性时操作 teacher.Sex, 当teacher无Sex时,操作的是 teacher.human.Sex
	tea.human.Sex = 200
	fmt.Println(tea)
	stu := student{
		Name:  "sss",
		Age:   13,
		human: human{Sex: 0},
	}
	fmt.Println(stu)
}

func A(per person) {
	per.Age = 12
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}

func C(per *person) {
	per.Age = 15
	fmt.Println("C", per)
}
