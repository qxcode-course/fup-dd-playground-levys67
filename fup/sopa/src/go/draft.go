package main
import "fmt"
func main() {
    qtd := 0
    fmt.Scan(&qtd)
    vetor := make([]int,qtd)
    for i := 0; i < len(vetor); i++{
        if i == 0 || i == 1 {
            vetor[i] = 1
        } else if i > 1 {
            vetor[i] = vetor[i-1] + vetor[i-2]
        }
        
    }
    fmt.Println(vetor[len(vetor)-1])
}