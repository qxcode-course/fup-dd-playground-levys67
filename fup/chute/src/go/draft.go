package main

import (
	"fmt"
	"math"
)
func main() {
    preçoreal := 0
    chute1 := 0
    chute2 := 0
    fmt.Scan(&preçoreal,&chute1,&chute2)
    dif1 := math.Abs(float64(preçoreal - chute1))
    dif2 := math.Abs(float64(preçoreal - chute2))
    if dif1 < dif2 {
        fmt.Println("primeiro")
    } else if dif2 < dif1 {
        fmt.Println("segundo")
    } else {
        fmt.Println("empate")
    }
}