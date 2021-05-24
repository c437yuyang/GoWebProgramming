package main

import "fmt"

type Person struct {
	name   string
	age    int
	weight int
}

type Kid struct {
	Person // 匿名字段,只写类型，无字段名
	hobby  string
}

type Adult struct {
	Person
	company string
}

func (p *Person) SayHi() {
	fmt.Printf("hi, i am a person, name is:%s, age:%d,weight:%d\n", p.name, p.age, p.weight)
}

func (p *Adult) SayHi() {
	fmt.Printf("hi, i am a adult, name is:%s, company is %s\n", p.name, p.company)
}

func main() {
	k1 := Kid{Person{name: "k1", age: 1, weight: 2}, "h1"}
	k1.SayHi()

	a1 := Adult{Person{name: "a1", age: 11, weight: 22}, "c1"}
	a1.SayHi()
}
