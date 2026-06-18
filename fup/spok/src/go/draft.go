package main

import (
	"fmt"
)
func main() {
    numero := 0
    fmt.Scan(&numero)
    texto := fmt.Sprint(numero)
    inverso := ""
    for i := len(texto) -1; i >= 0; i-- {
        inverso += string(texto[i])
    }
    if texto == inverso {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
} 