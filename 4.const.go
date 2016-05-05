package main

import "fmt"

var sss = "123"

const (
	a = "123"
	b = len(a)
	c
	d
)

const (
	aa, bb = 1, "222"
	cc, dd
)

const (
	aaa = 'A'
	bbb
	ccc = iota
	ddd
)

const (
	eee = iota
)

const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("a=" + a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)
	fmt.Println("cc=", cc)
	fmt.Println("dd=", dd)
	fmt.Println("aaa=", aaa)
	fmt.Println("bbb=", bbb)
	fmt.Println("ccc=", ccc)
	fmt.Println("ddd=", ddd)
	fmt.Println("eee=", eee)
	fmt.Println(^2)
	fmt.Println(!true)
	fmt.Println(1 << 10 << 10 >> 10)

	/*
		6: 0110
		11: 1011
		& 0010 = 2
		| 1111 = 15
		^ 1101 = 13
		&^ 0100 = 4
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	x := 0
	if x > 0 && (10/x) > 1 {
		fmt.Println("OK")
	}

	fmt.Println("KB=", KB)
	fmt.Println("MB=", MB)
	fmt.Println("GB=", GB)
	fmt.Println("TB=", TB)
	fmt.Println("PB=", PB)
	fmt.Println("EB=", EB)
	fmt.Println("ZB=", ZB)
	fmt.Println("YB=", YB)
}
