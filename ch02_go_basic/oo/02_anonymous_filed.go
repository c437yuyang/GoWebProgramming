package main

import "fmt"

type Human struct {
    name   string
    age    int
    weight int
}

type Student struct {
    Human      // 只写类型，无字段名
    speciality string
}

func main() {
    mike := Student{Human{
        name:   "mike",
        age:    1,
        weight: 2,
    }, "csgo"}

    fmt.Println(mike.name, mike.age, mike.weight) // 可以看到，可以直接访问匿名字段的成员，实现类似于继承
    mike.Human = Human{
        name:   "mik1",
        age:    11,
        weight: 22,
    }
    fmt.Println(mike.name, mike.age, mike.weight)

    // 当内外字段同名的时候，优先使用外面的，如果需要访问内部的，Bob.Human.phone这样来指明访问

}
