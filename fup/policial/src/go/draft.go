package main
import "fmt"

func criarvetor (num int) []int{
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func ordenarvetor(vetor []int) []int{
    for i := 0; i < len(vetor)-1; i++{
        for j := 0; j < len(vetor)-1; j++{
            if vetor[j] > vetor[j+1]{
                vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
            }
        }
    }
    return vetor
}

func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    ordenado := ordenarvetor(vetor)
    for i, v := range ordenado {
        if i != 0{
            fmt.Print(" ")
        }
        fmt.Printf("%d", v)
        if i == len(ordenado)-1{
            
            fmt.Print("\n")
        }
    }
}