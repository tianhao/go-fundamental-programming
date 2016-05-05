package main

import (
	"fmt"
	"time"
)

func Pingpong1(s []int) {
	s = append(s, 3)
}

func Pingpong2(s []int) []int {
	s = append(s, 3)
	return s
}

func main() {
	// case 1: 前篇地址修改后,无法修改原切片数据
	s1 := make([]int, 0)
	Pingpong1(s1)
	fmt.Println(s1)

	// case 2: fix case 1, 修改切片的函数返回修改后的数据, 原切片变量=函数返回的数据
	s2 := make([]int, 0)
	s2 = Pingpong2(s2)
	fmt.Println(s2)

	// case 3
	t1 := time.Now()
	fmt.Println(t1.Format(time.ANSIC))
	fmt.Println(t1.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t1.Format("Mon Jan _2 15:04:06 2006")) // 输出时间倒退了

	// case 4, 闭包问题,下面匿名函数输出的都是"c"
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
	// fix, 将闭包中的变量改成参数传入,会分别输出
	s = []string{"a", "b", "c"}
	fmt.Println(s)
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(3 * time.Second)
}
