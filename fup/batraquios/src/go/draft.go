package main
import "fmt"
func main() {
    tam1 := 0
    fmt.Scan(&tam1)
    vetor1 := make([]int, tam1)
    for i := 0; i < tam1; i++ {
        fmt.Scan(vetor1[i])
    }
    tam2 := 0
    fmt.Scan(&tam2)
    vetor2 := make([]int, tam2)
    for i := 0; i < tam2; i++{
        fmt.Scan(vetor2[i])
    }
    fmt.Println(vetor1)
    fmt.Println(vetor2)
}
