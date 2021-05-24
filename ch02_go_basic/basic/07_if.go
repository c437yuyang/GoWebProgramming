package main

import "fmt"

func fun1(x int) int {
    return x
}

func main() {
    {
        var i1 int = 1
        if i1 < 2 {
            fmt.Println("i1 lt 2, value:", i1)
        } else if i1 < 3 {
            fmt.Println("i1 lt 3, value:", i1)
        } else {
            fmt.Println("i1 gt 3, value:", i1)
        }

        if x := fun1(i1); x > 10 { // 可以套一个 声明语句
            fmt.Println(x)
        }
    }


}
