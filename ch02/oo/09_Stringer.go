package main

import (
	"fmt"
	"strconv"
)

type PrintableType struct {
	iField int
	sFiled string
}

func (p PrintableType) String() string {
	return "{\"iFiled\":\"" + strconv.Itoa(p.iField) + "\", \"sField\":\"" + p.sFiled + "\"}"
}

func main() {
	p := PrintableType{
		1, "one",
	}
	fmt.Println(p)
}
