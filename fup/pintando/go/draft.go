package main

import (
	"fmt"
	"math"
)
func main() {
    var a float32 = 0
    var b float32 = 0
    var c float32 = 0
    fmt.Scan(&a,&b,&c)
    var per float32 = (a +b +c) / 2
    var heron float32 = per * (per - a) * (per - b) * (per - c)
    var area float32 = float32(math.Sqrt(float64(heron))) 
    fmt.Printf("%.2f\n",area)
}
