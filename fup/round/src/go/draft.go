package main

import (
	"fmt"
	"math"
)
func main() {
    f := ""
    n := 0.00
    fmt.Scan(&f, &n)
    if f == "r" {
        fmt.Println(math.Round(float64(n)))
    }else if f == "c" {
        fmt.Println(math.Ceil(float64(n)))
    } else if f == "f" {
        fmt.Println(math.Floor(float64(n)))
    }
}
