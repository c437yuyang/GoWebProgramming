package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 让出时间片给别人，下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func main() {
	{
		go say("world")
		say("hello")
	}
}
