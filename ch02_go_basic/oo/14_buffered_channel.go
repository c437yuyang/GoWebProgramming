package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // channel 应该在生产者的地方关闭
}

func main() {
	// buffered channel 就是指定channel可以存储的元素数量，指定2，那么写入第三个的时候才会阻塞，直到其他的goroutine从channel读取数据，腾出空间
	{
		c := make(chan int, 2) // 3可以正常运行,1 会报错
		c <- 1
		c <- 2
		fmt.Println(<-c)
		fmt.Println(<-c)
	}

	{
		c := make(chan int, 10)
		go fibonacci(cap(c), c)
		for i := range c { // 可以不断的读取channel中的数据，直到channel被显式close掉
			fmt.Println(i)
		}
	}

}
