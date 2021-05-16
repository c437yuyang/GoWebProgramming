package main

import "fmt"

func main() {

    {
        var i1 int8 = 1
        var i2 int8 = 2
        c := i1 + i2
        fmt.Println(c)
        var i3 int = 3
        fmt.Println(i3)
    }

    {
        var c complex64 = 5 + 5i
        fmt.Println(c)

        var c1 complex128 = 5 + 5i
        fmt.Println(c1)
    }

    {
        var s1 string = "abcd"
        fmt.Println(s1)

        // 字符串不可变，修改会报错
        //s1[0] = 'c' //
        c := []byte(s1) // cast
        c[0] = 'c'
        s11 := string(c)
        fmt.Println(s11)
        // 或:
        s12 := "c" + s1[1:]
        fmt.Println(s12)

        // 可以用+连接
        var s2 = "efg"
        ss := s1 + s2
        fmt.Println(ss)

        // 多行字符串:
        s3 := `1231231
               hello
               hi`
        fmt.Println(s3)
    }

    fmt.Println("hello world")
}
