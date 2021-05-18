package main

import "fmt"

func main() {
	{
		// 可以存储任意类型的数值
		var a interface{}
		var i int = 5
		s := "hello world"
		a = i
		a = s
		fmt.Println(a)
	}
}
