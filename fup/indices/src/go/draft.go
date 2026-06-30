package main
import "fmt"

func criarvetor(num int) []int{
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func ordenarvetor (vetor []int) []int{
    indices := make([]int, 0)
    for i := 0; i < len(vetor)-1; i++ {
        for j := 0; j < len(vetor)-1-i; j++ {
            if vetor[j] > vetor[j+1] {
                vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
                indices = append(indices, j, j+1)
            }
        }
    }
    return indices
}


func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    fmt.Println(vetor)
    indices := ordenarvetor(vetor)
    fmt.Println(indices)
}