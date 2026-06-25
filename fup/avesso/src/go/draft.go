package main
import "fmt"

func criarvetor(tam int) []int {
    vetor := make([]int, tam)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}



func main() {
    testes := 0
    tam := 0
    grito := 0
    fmt.Scan(&testes, &tam, &grito)
    vetor := criarvetor(tam)
    fmt.Println(vetor)
}