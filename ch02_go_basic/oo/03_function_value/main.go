package main

import (
    "fmt"
    "image/color"
    "math"
)

type Point struct {
    X, Y float64
}

func (p *Point) Distance(q *Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(scale float64) {
    p.X *= scale
    p.Y *= scale
}

func (p *ColoredPoint) ScaleBy(scale float64) {
    p.X *= scale
    p.Y *= scale
}

type ColoredPoint struct {
    *Point
    Color color.RGBA
}

func main() {
    // 方法值
    {
        p := Point{1, 1}
        q := Point{5, 4}
        fmt.Println(p.Distance(&q))
        distanceP := p.Distance
        fmt.Println(distanceP(&q))
    }

    // 方法表达式
    {
        p := Point{1, 1}
        q := Point{5, 4}
        distanceP := Point.Distance
        fmt.Println(distanceP(p, &q))
    }
}
