package main

import "fmt"

func main() {
    const (
        a = iota
        b = iota
        c
    )
    const d = iota
    fmt.Println(a, b, c, d)
}
