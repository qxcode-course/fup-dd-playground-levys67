package main
import "fmt"

func criarvetor (num int) []int {
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func inverter (vetor []int) []int {
    inverso := make([]int,0)
    for i := len(vetor) - 1; i >= 0; i-- {
        inverso = append(inverso, vetor[i])
    }
    return inverso
}

func contaresforço(vetor []int) int {
    cont := 0
    for i := 0; i < len(vetor)-1; i++{
        if vetor[i] < vetor[i+1] {
            cont += vetor[i+1] - vetor[i]
        }
    }
    return cont
}



func main() {
    num := 0
    fmt.Scan(&num)
    trilha := criarvetor(num)
    trilhainvertida := inverter(trilha)
    contagem1 := contaresforço(trilha)
    contagem2 := contaresforço(trilhainvertida)
    if contagem1 < contagem2 {
        fmt.Println(contagem1)
    } else {
        fmt.Println(contagem2)
    }
}