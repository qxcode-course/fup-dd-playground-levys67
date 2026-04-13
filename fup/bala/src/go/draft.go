package main

import (
	"fmt"
	"math"
)
func main() {
    var a, b float64 = 0.00, 0.00
    var c, d float64 = 0.00, 0.00
    fmt.Scan(&a, &b, &c , &d)
    dac := c - a
    dbd := d - b
    dif := math.Sqrt(math.Pow(dac,2) + (math.Pow(dbd,2)))
    fmt.Printf("%.2f\n",dif)
}
