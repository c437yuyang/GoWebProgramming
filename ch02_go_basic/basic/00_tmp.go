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

type ColoredPoint struct {
    *Point
    Color color.RGBA
}

func main() {
    p := ColoredPoint{&Point{1, 1}, color.RGBA{}}
    q := ColoredPoint{&Point{5, 4}, color.RGBA{}}
    fmt.Println(p.Point, q.Point)
    q.Point = p.Point
    p.ScaleBy(2)
    fmt.Println(p.Point, q.Point)
}
