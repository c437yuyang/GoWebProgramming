package main

import "fmt"

func switchDemo(i int) {
    switch i {
    case 1:
        fmt.Println("i is 1")
    case 2, 3, 4:
        fmt.Println("i is 2,3,4")
        fallthrough // go 默认是break的，需要穿透的话加fallthr
    case 5:
        fmt.Println("i is 5")
    default:
        fmt.Println("not above")
    }
}

func main() {
    {
        i := 1
        switchDemo(i)
        i = 2
        switchDemo(2)
        i = 7
    }

}
