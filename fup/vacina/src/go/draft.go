package main
import "fmt"

func criarvetor (num int) []int {
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func contarmaior (vetor []int) int {
    maior := vetor[0]
    for i := 0; i < len(vetor); i++ {
        if maior < vetor[i] {
            maior = vetor[i]
        }
    }
    return maior
}


func main() {
    num := 0
    fmt.Scan(&num)
    vacinas := criarvetor(num)
    pessoas := criarvetor(num)
    contv := contarmaior(vacinas)
    contp := contarmaior(pessoas)
    fmt.Println(contv, contp)
}