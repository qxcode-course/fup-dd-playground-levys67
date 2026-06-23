package main
import "fmt"

func criarvetor(tam int) []int{
    vetor := make([]int,tam)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func separarforças(vetor []int) (int,int) {
    jedi := 0
    sith := 0
    for i, v := range vetor {
        if i <= (len(vetor) -1) / 2 {
            jedi += v
        } else {
            sith += v
        }
    }
    return jedi, sith
}


func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    jedi, sith := separarforças(vetor)
    if jedi > sith {
        fmt.Println("Jedi")
    } else if jedi == sith {
        fmt.Println("Empate")
    } else {
        fmt.Println("Sith")
    }
}