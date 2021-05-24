package main

import (
	"fmt"
	"reflect"
)

func main() {
	{
		human := Human{
			name:   "aa",
			age:    1,
			weight: 2,
		}
		t := reflect.TypeOf(human)
		fmt.Println(t)
	}

}
