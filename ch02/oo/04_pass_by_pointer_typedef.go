package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	GREEN
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // 类似于C里面的typedef

func (b *Box) SetColor(c Color) {
	//(*b).color = c  // 也可以写做这样,但是go不需要
	b.color = c
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		//(&bl[i]).SetColor(BLACK) // 检测到method的receiver是指针，自动帮转换了
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"white", "black", "blue", "red", "green"}
	return strings[c]
}

func main() {
	boxs := BoxList{
		Box{1, 2, 3, RED},
		Box{4, 5, 5, WHITE},
	}
	fmt.Println(boxs) // [{1 2 3 3} {4 5 5 0}]
	boxs.PaintItBlack()
	fmt.Println(boxs) // [{1 2 3 1} {4 5 5 1}], 可以看到成功修改了内容

	boxs[0].SetColor(GREEN)
	fmt.Println(boxs) // [{1 2 3 4} {4 5 5 1}], 修改了内容

	c := Color(WHITE) // Color实现了String方法
	fmt.Println(c)
}
