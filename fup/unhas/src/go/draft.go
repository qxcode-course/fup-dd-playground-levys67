 package main
import "fmt"

func criarvetor(qtd int) []int{
    vetor := make([]int, qtd)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func igualazero (vetor []int) bool {
    for i := 0; i < len(vetor); i++{
        if vetor[i] != 0{
            return false
        }
    }
    return true
}


func main() {
    qtd := 0
    fmt.Scan(&qtd)
    vetor := criarvetor(qtd)
    if igualazero(vetor) {
        fmt.Println(0)
    } else {
        for i := 0; i < len(vetor); i++{
        if i == 0 && vetor[i] == 0 {
            continue
        }
        fmt.Print(vetor[i])
    } 
    fmt.Println()

    }

}