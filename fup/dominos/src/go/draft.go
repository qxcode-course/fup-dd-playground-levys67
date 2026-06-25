package main
import "fmt"

func criarvetor(num int) []int {
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func conferirsequencia(vetor []int) bool {
    for i := 0; i < len(vetor)-1; i++{
        if vetor[i] > vetor[i+1] {
            return false
        }
    }
    return true
}


func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    ehcrescente := conferirsequencia(vetor)
    if !ehcrescente {
        fmt.Println("precisa de ajuste")
    } else {
        fmt.Println("ok")
    }
}