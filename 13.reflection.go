package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manger struct {
	User
	title string
}

func (u User) Hello(name string) (int, int) {
	fmt.Println("Hello ", name, ",", "my name is ", u.Name)
	return 10, 20
}

func main() {
	u := User{1, "OK", 12}
	Info(u)
	// Info(&u)

	m := Manger{User: User{1, "OK", 12}, title: "jack"}
	t := reflect.TypeOf(m)
	fmt.Println("---------")
	fmt.Printf("Field(0): %#v\n", t.Field(0))
	fmt.Println("---------")
	fmt.Printf("Field(1): %#v\n", t.Field(1))
	fmt.Println("---------")
	fmt.Printf("FieldByIndex[0,1]: %#v\n", t.FieldByIndex([]int{0, 1}))

	// 通过反射修改数据
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	uu := User{1, "OK", 12}
	Set(&uu)
	fmt.Println(uu)

	uu.Hello("joe")
}

func Info(o interface{}) {
	//
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("类型错误")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	// 取字段并调用它
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// 取方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

	// 通过反射来调用对应方法
	md := v.MethodByName("Hello")                   // 获得方法引用
	args := []reflect.Value{reflect.ValueOf("joe")} // 设置参数
	ret := md.Call(args)                            // 调用方法
	fmt.Println("返回值:", ret)
	fmt.Println("返回值:", reflect.ValueOf(ret))
	fmt.Println("返回值类型:", reflect.TypeOf(ret))
	l := len(ret)
	fmt.Println("返回值长度", len(ret))
	fmt.Println("分别打印返回值:")
	for i := 0; i < l; i++ {
		fmt.Println(ret[i])
	}

}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	// 非指针 或者不能设置则退出
	if v.Kind() != reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不为指针或者不能设置值")
		return
	} else {
		v = v.Elem()
	}

	// 通过字段名称取字段
	f := v.FieldByName("Name")
	// 无效字段忽略
	if !f.IsValid() {
		fmt.Println("BAD")
	}
	// 如果字段是字符串类型,则设置新值
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
