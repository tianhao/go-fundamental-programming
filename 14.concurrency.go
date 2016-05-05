package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	// case 1: bad
	printCase(1)
	go Go()
	time.Sleep(2 * time.Second)

	// case 2
	printCase(2)
	c2 := make(chan bool)
	go func() {
		fmt.Println("case 2: Go Go Go!!")
		c2 <- true
	}()
	<-c2

	// case 3
	printCase(3)
	c3 := make(chan bool)
	go func() {
		fmt.Println("case 3: Go Go Go!!!")
		c3 <- true
		close(c3)
	}()
	// <-c
	for v := range c3 {
		fmt.Println("case 3, receive:", v)
	}

	// case 4: 有缓存,先取再读, BAD
	// printCase(4)
	// c4 := make(chan bool, 1)
	// <-c4
	// go func() {
	// 	fmt.Println("case 4: Go Go Go!!!")
	// 	c4 <- true
	// }()

	// case 5: 有缓存时, 主线程存, 子线程取, 主线程先退出, 子线程无效, BAD --> 除非无缓存
	printCase(5)
	c5 := make(chan bool, 1)
	go func() {
		fmt.Println("case 5: Go Go Go!!!")
		<-c5
	}()
	c5 <- true

	// case 6: 使用并发, 在最后一个并发中往channel中写入值通知主线程结束,
	// BAD: 最后一个启动的并发不一定是最后结束, 多核cpu调度不是按启动顺序调度的
	printCase(6)
	runtime.GOMAXPROCS(runtime.NumCPU()) // 经测试,在mac中不需要设置这个参数也能使用多核
	c6 := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go2(c6, i)
	}
	<-c6

	// case 7: fix case 6, 所有启动的并发都往channel中写入执行结束通知, 主线程从channel中读取10次
	printCase(7)
	c7 := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go3(c7, i)
	}

	for i := 0; i < 10; i++ {
		<-c7
	}

	// case 8: fix case 6, 通过sync.WaitGroup方法等待所有任务结束
	printCase(8)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go4(&wg, i) // 注意这里传递的是指针,非值拷贝
	}
	wg.Wait()

	// case 9: select
	c9x, c9y := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		a, b := false, false
		for {
			select {
			case v, ok := <-c9x:
				if !ok {
					fmt.Println("c1")
					o <- true
					break
					if !a { //此方法无效, 因为即使c9x没有数据, select默认读到默认值0,所以不能在这里取消写
						a = true
						o <- true
						break
					}
				}
				fmt.Println("c9x", v)
			case v, ok := <-c9y:
				if !ok {
					fmt.Println("c2")
					o <- true
					break
					if !b { // 方法无效, 见上面c9x说明
						b = true
						o <- true
						break
					}
				}
				fmt.Println("c9y", v)
			}
		}
	}()
	c9x <- 1
	c9y <- "hi"
	c9x <- 3
	c9y <- "hello"
	close(c9x)
	close(c9y)

	for i := 0; i < 2; i++ {
		<-o
	}

	// case 10: 空select{}会阻塞?, 可用于阻塞main函数不退出?
	// 经测试,最近版本(1.6)不能使用此方法,报错: goroutine 1 [select (no cases)]:
	printCase(10)
	// select {} // error

	// case 11: 为select设置超时
	printCase(11)
	c11 := make(chan bool)
	select {
	case v := <-c11:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

	// case 12: 作业
	printCase(12)
	c12 := make(chan string)
	go func() { //Pingpong
		i := 0
		for {
			fmt.Println(<-c12)
			c12 <- fmt.Sprintf("From Pingpong: Hi, #%d", i)
			i++
		}
	}()

	for i := 0; i < 10; i++ {
		c12 <- fmt.Sprintf("From main: Hello, #%d", i)
		fmt.Println(<-c12)
	}
}

func printCase(caseN int) {
	fmt.Println("-------------- case", caseN, "-------------")
}

func Go() {
	fmt.Println("Go Go Go!")
}

// 最后一个任务才写入channel, BAD
func Go2(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("Go2", index, a)

	if index == 9 {
		c <- true
	}
}

// 每个任务都写入channel方案
func Go3(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("Go3", index, a)
	c <- true
}

// sync解决方案, 注意WaitGroup是传递的指针
func Go4(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("Go4", index, a)
	wg.Done()
}
