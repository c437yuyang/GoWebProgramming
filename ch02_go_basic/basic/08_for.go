package main

import "fmt"

func getDayOfWeek(y int, m int, d int) int {
    if m <= 2 {
        m += 12
        y--
    }
    return (d + 2*m + 3*(m+1)/5 + y + y/4 - y/100 + y/400 + 1) % 7
}

func main() {
    d := 1
    y := 2021
    {
        // 第一种 for exp1;exp2;exp3 {}
        for m := 1; m < 13; m++ {
            fmt.Println(m, getDayOfWeek(y, m, d))
            if m%2 == 0 {
                continue
            }
        }

        // go 没有“,”操作符,因此需要使用平行赋值当有两个需要赋值的时候
        for m, m1 := 1, 2; m < 13; m++ { // 但是后面的表达式怎么处理，比如要m1++???
            fmt.Println(m, getDayOfWeek(y, m, d))
            fmt.Println(m, getDayOfWeek(y, m1, d))
        }

        // 可以省略
        m := 1
        for m < 13 {
            fmt.Println(m, getDayOfWeek(y, m, d))
            if m == 1 {
                break
            }
        }
    }
    {
        m1 := map[string]int{"1": 1, "2": 2}
        // 第二种配合range读取slice和map
        for k, v := range m1 {
            fmt.Println(k, ":", v)
        }

        for _, v := range m1 { // 可以丢弃
            fmt.Println(v)
        }

        s1 := []int{1, 2, 3}
        for idx, num := range s1 { // 遍历slice, 是idx和ele
            fmt.Println(idx, num)
        }
    }

}
