package main

import (
    "fmt"
    "math"
)

type shape interface {
    area() float64
    perimeter() float64
}

type circle struct {
    radius float64
}

type rectangle struct {
    width, height float64
}

func(c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func(c circle) perimeter() float64 {
    return  2* math.Pi * c.radius
}

func(r rectangle) area() float64 {
    return r.width * r.height
}

func(r rectangle) perimeter() float64 {
    return 2 * (r.width + r.height)
}

func measure(s shape) {
    fmt.Println("Area:", s.area())
    fmt.Println("Perimeter:", s.perimeter())
}

func main() {
    r := rectangle{5, 5}
    c := circle{5}
    measure(r)
    measure(c)
}
