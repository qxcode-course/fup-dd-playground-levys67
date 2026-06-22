package main
import "fmt"

func criarvetor(num int) []int{
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    fmt.Print("[")
    for i := len(vetor) -1; i >= 0; i--{
        if i == len(vetor) -1 {
            fmt.Print(" ")
        }
        fmt.Print(vetor[i], " ")
    }
    fmt.Println("]")
}