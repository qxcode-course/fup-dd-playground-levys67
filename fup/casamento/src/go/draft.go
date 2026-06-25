package main
import "fmt"

func contarmenor (vetor []int) int{
    menor := vetor[0]
    for _ , v := range vetor{
        if v < menor {
            menor = v
        }
    }
    return menor
}

func contarmaior (vetor []int) int {
    maior := vetor[0]
    for _ , v := range vetor {
        if v > maior {
            maior = v
        }
    }
    return maior
}



func main() {
    vetor := make([]int, 5)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    menor := contarmenor(vetor)
    maior := contarmaior(vetor)
    fmt.Println(maior + menor)
}