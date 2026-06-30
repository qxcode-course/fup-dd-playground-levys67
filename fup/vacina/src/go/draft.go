package main
import "fmt"

func criarvetor (num int) []int {
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func ordenarvetor(vetor[]int) []int{
    for i := 0; i < len(vetor)-1; i++{
        for j := 0; j < len(vetor)-1-i ; j++{
            if vetor[j] > vetor[j+1] {
                vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
            }
        }
    }
    return vetor
}

func compararvetores (vetor1, vetor2 []int) bool {
    for i := 0; i < len(vetor1); i++{
        if vetor1[i] < vetor2[i]{
            return false
        }
    }
    return true
}


func main() {
    num := 0
    fmt.Scan(&num)
    vacinas := criarvetor(num)
    pessoas := criarvetor(num)
    ordv := ordenarvetor(vacinas)
    ordp := ordenarvetor(pessoas)
    relatorio := compararvetores(ordv, ordp)
    if !relatorio{
        fmt.Println("No")
    } else {
        fmt.Println("Yes")
    }
}