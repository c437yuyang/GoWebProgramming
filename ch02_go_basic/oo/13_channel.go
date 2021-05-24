package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, num := range a {
		sum += num
	}
	c <- sum // send sum to channel c
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	{
		// goroutine 之间通过channel来做同步
		c := make(chan int)
		go sum(a[:len(a)/2], c)
		go sum(a[len(a)/2:], c)
		x, y := <-c, <-c // channel的发送和接收都是阻塞的，不需要显式的lock
		// 所以这里输出可能是40 15 55 也可能是 15 40 55，就看哪个goroutine先完成
		fmt.Println(x, y, x+y)
	}

	{
		c1, c2 := make(chan int), make(chan int)
		go sum(a[:len(a)/2], c1)
		go sum(a[len(a)/2:], c2)
		x, y := <-c1, <-c2
		fmt.Println(x, y, x+y) // 这里输出就一定是 15 40 55了
	}
}
