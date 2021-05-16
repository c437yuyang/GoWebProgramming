package main

import "fmt"

func printNumber(i int) {
    defer fmt.Println("fun1 is returning")          // 一个只能一行语句
    defer fmt.Println("fun1 is about to returning") // 可以多个，逆序执行
    for i := 0; i < 2; i++ {
        defer fmt.Println(i)
    }
    defer func() {
        fmt.Println("defer func")
    }()
    fmt.Println(i)
}

func main() {
    printNumber(10)
}
