package main

import (
    "fmt"
    "strings"
)

func testPanic(str string) {
    if str == "a" {
        panic("str is a")
    }
}

func testRecover() (bret bool) {
    defer func() {
        if x := recover(); x != nil {
            bret = true // 配合defer, 这里可以直接修改返回值
        }
    }()
    testPanic("a")
    return false
}

func main() {
    fmt.Println(testRecover()) // true

    res := strings.Map(func(r rune) rune {
        return r + 1
    }, "HAL-9000")
    fmt.Println(res)
}
