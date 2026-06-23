package main
import "fmt"

func criarvetor (tam int) []int{
    vetor := make([]int, tam)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func fazendoparkour(vetor []int) int{
    pulos := 0
    for i := 0; i < len(vetor)-1; i++{
        if vetor[i+1] - vetor[i] > 1 || vetor[i+1] - vetor[i] < -1 {
            pulos ++
        }
        
    }
    return pulos
}


func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    pulos := fazendoparkour(vetor)
    fmt.Println(pulos)
}