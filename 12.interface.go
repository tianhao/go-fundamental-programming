package main

import (
	"fmt"
)

// type USB interface {
// 	Name() string
// 	Connect()
// }

type empty interface {
}

type Connecter interface {
	Connect()
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnecter struct {
	name string
}

type TVConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func (tv TVConnecter) Connect() {
	fmt.Println("Connect:", tv.name)
}

// func Disconnect(usb USB) { // 这里要求的是USB类型, 而USB是interface类型
// 	fmt.Println("Disconnect:", pc.Name())
// }

// func Disconnect(usb empty) { // 这里要求的是USB类型, 而USB是interface类型
// 	if pc, ok := usb.(PhoneConnecter); ok {
// 		fmt.Println("Disconnect:", pc.Name())
// 		return
// 	}
// 	fmt.Println("Unknown device")
// }

func Disconnect(usb interface{}) { // 这里要求的是USB类型, 而USB是interface类型
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect:", v.Name())
	default:
		fmt.Println("Unknown device.")
	}
}

func main() {
	// var usb USB
	// usb = PhoneConnecter{"PhoneConnector"}
	pc := PhoneConnecter{"PhoneConnector"}
	pc.Connect()
	Disconnect(pc)

	other := 1
	Disconnect(other)

	var con Connecter
	con = Connecter(pc)
	con.Connect()

	pc.name = "pc"
	// 注意观察下面输出
	// 发现name还是"PhoneConnector", 不是上面修改后的"pc", 说明con是pc的复制品, 不是引用
	con.Connect()

	// con.Name() // error 转换为Connector类型后就不能非Connector的方法了
	// 注意观察下面语句输出
	// 这里con.(type)还是识别类型为PhoneConnector类型,因为数据类型没变, 上面变得是接口类型
	Disconnect(con)

	// tv := TVConnecter{"TVConnector"}
	// var usb USB
	// usb = USB(tv) // error: 由于tv没有实现USB接口, 所以不能强制转换
	// usb.Connect()

	var a interface{}
	fmt.Println(a == nil)
	var p *int = nil
	a = p
	fmt.Println(a == nil) // a虽然指向空指针, 但依然不是nil

}
