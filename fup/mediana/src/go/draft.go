package main
import "fmt"

func criarvetor( num int )[]float64 {
    vetor := make([]float64, num)
    for i := 0; i<len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func ordenarvetor(vetor []float64) []float64 {
    for i := 0; i < len(vetor)-1; i++ {
        for j := 0; j < len(vetor)-1-i; j++{
            if vetor[j] > vetor[j+1] {
                vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
            }
        }
    }
    return vetor
}

func mediana (vetor []float64) float64{
    if len(vetor) % 2 == 0{
        media := vetor[len(vetor)/2] + vetor[(len(vetor)/2)-1]
        return media / 2
    }
    return vetor [len(vetor)/2]
}


func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    notas := ordenarvetor(vetor)
    mediana := mediana(notas)
    fmt.Printf("%.1f\n", mediana)
}