package main

import (
	"fmt"
	"math"
)
func main() {
    a := 0.0
    b := 0.0
    c := 0.0
    fmt.Scan(&a, &b, &c)
    delta := b*b - 4*a*c
    if  delta < 0 {
        fmt.Println("nao ha raiz real")
        return
    }
    x1 := (-b + math.Sqrt(delta)) / (2*a)
    x2 := (-b - math.Sqrt(delta)) / (2*a)
    if delta == 0 {
        fmt.Printf("%.2f\n", x1)
    } else if delta > 0 {
        fmt.Printf("%.2f\n%.2f\n", x1, x2)
    }
}
