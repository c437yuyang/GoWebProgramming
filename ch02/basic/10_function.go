package main

import "fmt"

func max(i1 int, i2 int) int { // func funcName(input1 type1, input2 type2) (output type3, output type4) {return v1,v2}
    if i1 > i2 {
        return i1
    } else {
        return i2
    }
}

func echo(i1 int, i2 int) (int, int) {
    return i1, i2 // 可以多返回值
}

func print(i int) {
    fmt.Println(i)
}

// SumAndProduct 官方建议如果是导出函数，建议使用命名返回值
func SumAndProduct(A, B int) (SumRes int, MulRes int) {
    SumRes = A + B
    MulRes = A * B
    return
}

// 变参
func printIntegers(numbers ...int) {
    for idx, num := range numbers { // 实际上在函数内,numbers就是个int的slice
        fmt.Printf("numbers[%d]:%d,", idx, num)
    }
    fmt.Println()
}

// 传指针
func incr(num *int) {
    *num += 1
}

func main() {
    {
        fmt.Println(SumAndProduct(2, 5))
        printIntegers(1, 2, 3, 4, 5)
    }
    {
        num := 1
        incr(&num)
        fmt.Println(num)
    }

}
