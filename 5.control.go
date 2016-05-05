package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")

	a := 1
	var p *int = &a
	fmt.Println("*p=", *p)
	// a := a++ // error
	a++
	// ++a // error
	a = 10
	if a, b := 1, 2; a > 0 {
		fmt.Println("a=", a, ",b=", b)
	}
	fmt.Println("a=", a)

	a = 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println("in for, a=", a)
	}
	fmt.Println("Over")
	a = 1
	for a <= 3 {
		a++
		fmt.Println("in for, a=", a)
	}
	fmt.Println("Over")

	for i := 0; i < 3; i++ {
		a++
		fmt.Println("in for, a=", a)
	}
	fmt.Println("Over")

	fmt.Println("Switch 1")
	a = 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("Switch 2")
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("Switch 3")
	switch s := 1; {
	case s >= 0:
		fmt.Println("s=0")
		fallthrough
	case s >= 1:
		fmt.Println("s=1")
	default:
		fmt.Println("None")
	}
	// fmt.Println("s=", s) // error, s undefined
BREAK_LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break BREAK_LABEL
			}
		}
	}
	fmt.Println("break to BREAK_LABEL")

CONTINUE_LABEL:
	for i := 0; i < 10; i++ {
		for {
			continue CONTINUE_LABEL
			fmt.Println("i=", i)
		}
	}
	fmt.Println("continue to CONTINUE_LABEL")

GOTO_LABEL:
	for i := 0; i < 10; i++ {
		for {
			goto GOTO_LABEL
			fmt.Println("i=", i)
		}
	}

	fmt.Println("goto to GOTO_LABEL")
}
