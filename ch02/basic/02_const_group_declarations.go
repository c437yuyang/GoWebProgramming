package main

import (
    "fmt"
    "os"
)

// 一些声明可以写在一起
const (
    ci      = 100
    cpi     = 3.14
    cprefix = "go_"
)

const (
    a = 1
    b
    c
)

func main() {
    // 分组声明
    var (
        i1      int
        pi1     float32
        prefix1 string
    )
    fmt.Println(i1, pi1, prefix1)
    fmt.Println(os.Args)
    fmt.Println(a, b, c)
}
