package main

import "fmt"

func main() {
    {
        var arr [2]string
        arr[0] = "abc"
        fmt.Println(arr)

        var arr1 [2]int
        arr1[0] = 0
        fmt.Println(arr1) // 不赋值默认是0

        arr2 := [4]int{1, 2, 3} // 不够的填0
        fmt.Println(arr2)       // [1 2 3 0]

        arr3 := [...]int{1, 2, 3} // 可以自动推断长度
        fmt.Println(arr3)
    }
    {
        var arr1 [2]int
        arr1[0] = 0
        arr1[1] = 1
        arr2 := arr1

        arr2[0] = 1
        fmt.Println(arr1, arr2) // [0 1] [1 1], 可以看到是拷贝的
    }

    {
        // 二维数组
        a1 := [2][3]int{[3]int{1, 2, 3}, [3]int{2, 3}}
        fmt.Println(a1)

        a2 := [...][3]int{{1, 2, 3}, {2, 3, 4}} // 1. 可以省略内部数组的类型声明 2.也可以用自动推导（外层可以，内层不行）
        fmt.Println(a2)
    }
}
