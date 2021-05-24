package main

import "fmt"

func main() {
    {
        var m1 map[string]int
        m1 = make(map[string]int) // 使用前需要初始化
        m1["one"] = 1
        m1["two"] = 2
        fmt.Println(m1, m1["one"], m1["twp"]) // map[one:1 two:2] 1 0, 找不到的返回0???

        m2 := make(map[string]int) // 也可直接make一个出来
        fmt.Println(m2)

        m3 := map[string]int{"1": 1, "2": 2} // 可以直接初始化
        fmt.Println(m3)
    }

    {
        // 使用
        m1 := map[string]int{"1": 1, "2": 2}
        fmt.Println(m1, m1["1"], m1["11"]) // map[1:1 2:2] 1 0, 找不到的返回0???
        v, ok := m1["1"]                   // v + ok的方式
        fmt.Println(v, ok)
        v, ok = m1["11"]
        fmt.Println(v, ok)

        delete(m1, "1") // 使用delete 删除key
        fmt.Println(m1)

        m2 := m1
        m2["1"] = 11
        fmt.Println(m1, m2) // map[1:11 2:2] map[1:11 2:2], 可以看到都被修改了
    }
}
