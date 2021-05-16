package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {

    {
        var p1 person
        p1.name = "name"
        p1.age = 10

        fmt.Println(p1) // {name 10}

        // 或
        p2 := person{"name1", 11}
        fmt.Println(p2)

        // 或
        p3 := person{name: "name2", age: 12}
        fmt.Println(p3)
    }

}
