package main

import (
	"fmt"
	"sort"
)

func main() {
	// 单层map
	// var m map[int]string = make(map[int]string)
	m := make(map[int]string)
	m[1] = "OK"
	fmt.Println(m)
	a := m[1]
	fmt.Println(1, a)
	b := m[2]
	fmt.Println(2, b)
	// c := &m[1] //  cannot take the address of m[1]
	delete(m, 1)
	fmt.Println(m)
	// fmt.Println(1, c)

	// 两层map, 第二层使用前需要初始化
	var mm map[int]map[int]string
	mm = make(map[int]map[int]string)
	fmt.Println(mm)
	// mm[1][1] = "OK" // error: panic: assignment to entry in nil map
	mm[1] = make(map[int]string) // 初始化 mm[1]
	mm[1][1] = "OK"              // success
	// mm[2][1] = "OK"              // error: panic: assignment to entry in nil map
	fmt.Println(mm[2][1])
	a, ok := mm[2][1]
	if !ok {
		mm[2] = make(map[int]string)
	}
	mm[2][1] = "GOOD"
	a, ok = mm[2][1]
	fmt.Println(a, ok)

	// map 迭代
	// for i, v := range slice {
	// 	slice[i] = xxx
	// }
	// for k, v := range mm {
	// 	mm[k] = xxx
	// }

	// 对拷贝操作 --> 无效操作
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)

	// 对实际数据操作 --> 有效操作
	sm2 := make([]map[int]string, 5)
	for i := range sm2 {
		sm2[i] = make(map[int]string, 1)
		sm2[i][1] = "OK"
		fmt.Println(sm2[i])
	}
	fmt.Println(sm2)

	// 对map的key做排序的方法: 使用slice存储key, 对slice排序
	mmm := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	sss := make([]int, len(mmm))
	i := 0
	for k, _ := range mmm {
		sss[i] = k
		i++
	}
	sort.Ints(sss)
	fmt.Println(sss)

	x := map[int]string{1: "a", 2: "b", 3: "c"}
	y := make(map[string]int)
	for k, v := range x {
		y[v] = k
	}
	fmt.Println(x)
	fmt.Println(y)
}
