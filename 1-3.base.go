// 当前程序包名
package main

// 导入其它包
import "fmt"
import "math"

// 全局变量声明不能省略var关键字
var PI = 3.14

// 全局变量声明与复制
var name = "gopher"

// 一般类型声明
type newType int

// 结构类型声明
type gopher struct{}

// 接口类型声明
type golang interface{}

type byte int8
type rune int32
type 文本 string

// main函数声明
func main() {
	var a [1]bool
	var b 文本
	b = "中文类型名"
	var c int = 1
	d := false
	var x, y, z = 1, 2, 3
	xx, _, zz := 3, 2, 1
	var aa float32 = 100.1
	bb := int(aa)
	// cc := bool(aa)
	fmt.Println(bb)
	fmt.Println("Hello Go!")
	fmt.Println(a)
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(xx)
	fmt.Println(zz)
}
