package main

import "fmt"

type numberFilter func(int) bool // 声明函数类型
func isOdd(i int) bool {
    return i%2 == 1
}
func isEvent(i int) bool {
    return i%2 == 0
}

func printIf(slice []int, f numberFilter) {
    for _, num := range slice {
        if f(num) {
            fmt.Print(num, ",")
        }
    }
}

func main() {
    s1 := []int{1, 2, 3}
    printIf(s1, isOdd)
}
