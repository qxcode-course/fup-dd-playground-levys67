package main

import (
	"fmt"
	"math"
)
func main() {
    a, b := 0.00, 0.00
    c, d := 0.00, 0.00
    fmt.Scan(&a,&b,&c,&d)
    dist1 := c - a
    dist2 := d - b
    som := math.Sqrt(math.Pow(dist1,2,) + math.Pow(dist2,2))
    fmt.Printf("%0.2f\n",som)
}
