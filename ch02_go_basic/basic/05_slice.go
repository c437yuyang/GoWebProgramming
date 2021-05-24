package main

import "fmt"

func main() {
    {
        var s1 []int
        s2 := []int{1, 2}

        fmt.Println(s1, s2)
    }

    {
        // 可以从已存在的数组或slice再次声明
        s1 := [4]int{1, 2, 3, 4}
        s2 := s1[1:3]   // 前闭后开 [)
        fmt.Println(s2) // [2 3]

        fmt.Println(s1[:3]) // 不写的话,pos1默认从0开始, pos2默认是n
        fmt.Println(s1[:])
    }

    {
        // 引用类型，修改会改到原来的值
        arr := [4]int{1, 2, 3, 4}
        slice := arr[1:3]
        slice[1] = 22
        fmt.Println(arr, slice) // [1 2 22 4] [2 22]
    }

    {
        // 几个有用的内置函数
        arr := [4]int{1, 2, 3, 4}
        s1 := arr[1:3]
        fmt.Println(len(s1))
        fmt.Println(cap(s1))
        s2 := append(s1, 1)
        s2[0] = 22
        fmt.Println(s1, s2) // [22 3] [22 3 1]

        var s3 []int
        n := copy(s2, s3) // 返回拷贝的个数
        //s3[0] = 222
        fmt.Println(n, s2, s3)
    }

}
