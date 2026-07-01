package main
import "fmt"

func criarvetor (qtd int) []int {
    vetor := make([]int, qtd)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func percorrer (vetor []int,limA, limB int) int {
    cont := 0
    for _, v := range vetor {
        if v >= limA && v <= limB {
            cont ++
        }
    }
    return cont
}


func main() {
    qtd := 0
    limA := 0
    limB := 0
    fmt.Scan(&qtd,&limA,&limB)
    vetor := criarvetor(qtd)
    contagem := percorrer(vetor,limA,limB)
    fmt.Println(contagem)
}
