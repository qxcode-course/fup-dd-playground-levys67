package main

import (
	"fmt"
	"math"
)
func main() {
    var a float64 = 0
    var b float64 = 0
    fmt.Scan(&a,&b)
    var c float64 = a - b
    c = math.Abs(c)
    fmt.Println(c)
}
