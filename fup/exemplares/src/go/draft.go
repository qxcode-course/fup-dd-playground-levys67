package main

import (
	"fmt"
	"sort"
)

func criarvetor(tam int) []int{
    vetor := make([]int, tam)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func existe(vetor []int, numero int) bool{
    for _, v := range vetor {
        if v == numero {
            return true
        }
    }
    return false
}

func contaranimais (vetor []int) []int {
    animais := make([]int, 0)
    for i := 0; i < len(vetor); i++{
        if !existe(animais, vetor[i]) {
            animais = append(animais, vetor[i])
        }
    }
    return animais
}


func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    animais := contaranimais(vetor)
    sort.Ints(animais)
    for i, v := range animais {
        if i != 0{
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println("")
}