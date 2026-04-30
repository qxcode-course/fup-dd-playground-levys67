package main

import (
	"fmt"
	"math"
)
func main() {
    chicobento := 0
    cebolinha := 0
    quantidade := 0
    fmt.Scan(&chicobento, &cebolinha, &quantidade)
    animais := make([]string, quantidade)
    for i := 0; i < quantidade; i++ {
        fmt.Scan(&animais[i])
    }
    total := 0
    c := 0
    g := 0
    v := 0
    for _, valor := range animais {
        if valor == "c"{
            c += 4
        } else if valor == "g" {
            g += 2
        } else if valor == "v" {
            v += 4
        }
    }
    total += c + g + v
    fmt.Println(total)
    dif1 := math.Abs(float64(chicobento) - float64(total))
    dif2 := math.Abs(float64(cebolinha) - float64(total))
    if dif1 > dif2 {
        fmt.Println("Cebolinha")
    } else if dif2 > dif1 {
        fmt.Println("Chico Bento")
    } else {
        fmt.Println("empate")
    }
}
