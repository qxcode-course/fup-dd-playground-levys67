package main
import "fmt"

func criarvetor(num int) []int {
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func contarparedes(vetor []int) int {
    cont := 0
    maior := 0
    for i := 0; i < len(vetor) -1; i++{
        if vetor[i] > maior {
            cont ++
            maior = vetor[i]
        }
        if maior < vetor[i+1] {
            cont ++
            maior = vetor[i+1]
        }
        
    }
    return cont
} 


func main() {
    num := 0
    fmt.Scan(&num)
    paredes := criarvetor(num)
    contagem := contarparedes(paredes)
    fmt.Println(contagem)
}