package main
import "fmt"

func criarvetor(tam int) []int {
    lista := make([]int, tam)
    for i := 0; i < tam; i++ {
        fmt.Scan(&lista[i])
    }
    return lista
}

func estacontido (vetor []int, elem int) bool {
    for i := 0; i < len(vetor); i++ {
        if elem == vetor[i]{
            return true
        }
    }
    return false
}



func compararvetores (vetor1, vetor2 []int) bool{
    for i := 0; i < len(vetor1); i++ {
        if !estacontido(vetor2, vetor1[i]) {
            return false
        }
    }
    return true
}


func main() {
    tam1 := 0
    tam2 := 0
    fmt.Scan(&tam1)
    vet1 := criarvetor(tam1)
    fmt.Scan(&tam2)
    vet2 := criarvetor(tam2)
    contido := compararvetores(vet1, vet2)
    if contido {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
