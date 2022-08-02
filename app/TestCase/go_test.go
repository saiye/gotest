package testcase

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//test1--协程执行测试,执行代码，你会看到输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行：
func Test1(t *testing.T) {
	go say("world")
	say("hello")
}

//输出字符串5次
func say(s string) {
	for i := 0; i < 5; i++ {
		//间隔一段时间
		//1/10秒
		time.Sleep(100 * time.Millisecond)
		//每秒
		//time.Sleep(time.Second)
		fmt.Println(s)
	}
}

//协程执行测试，分步计算结果
func Test2(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	l := len(s)
	fmt.Println("len：", l)
	go sum(s[:l/2], c)
	go sum(s[l/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

/**
通道（channel）是用来传递数据的一个数据结构。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据,并把值赋给 v
*/
func Test3(t *testing.T) {
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println("read:", <-ch, "--", <-ch)
	//再次接收，死锁，all goroutines are asleep - deadlock!
	//fmt.Println(<-ch)
}

func Test4(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//如果不CLOSE,会导致错误
	//fatal error: all goroutines are asleep - deadlock!
	//致命错误:所有的goroutines是睡眠-死锁!
	close(c)
}

func Test5(t *testing.T) {
	/*	f:= func(v int32) {
		fmt.Println(v)
	}*/
	var i int32 = 0
	for ; i < 10; i++ {
		go myPrint(i)
	}
	//让出时间片，让别的线程执行先
	runtime.Gosched()
	fmt.Println("main")
}
func myPrint(i int32) {
	fmt.Println(i)
}

func Test6(t *testing.T) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
}
