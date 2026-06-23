 package main
import "fmt"

func criarvetor (num int) []int{
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func contarnaoadjacentes(vetor []int) int{
    cont := 0
    for i := 0; i < len(vetor); i++{
        if i == 0 && vetor[i] == 0 && vetor[i+1] == 0 {
            cont ++
        } else if i == len(vetor) -1 && vetor[i] == 0 && vetor[i-1] == 0 {
            cont ++
        } else if vetor[i] == 0 && vetor[i+1] == 0 && vetor[i-1] == 0 {
            cont ++
        }
    }
    return cont
}


func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    contagem := contarnaoadjacentes(vetor)
    fmt.Println(contagem)
}