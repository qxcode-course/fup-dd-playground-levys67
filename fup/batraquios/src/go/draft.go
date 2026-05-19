package main
import "fmt"

func criarvetor(tam int) []int {
    lista := make([]int, tam)
    for i := 0; i < tam; i++ {
        fmt.Scan(&lista[i])
    }
    return lista
}

func main() {
    tam1 := 0
    tam2 := 0
    fmt.Scan(&tam1)
    vet1 := criarvetor(tam1)
    fmt.Scan(&tam2)
    vet2 := criarvetor(tam2)
    //fmt.Print(vet1, vet2)
    contido := true
    for i := 0; i < tam1; i++ {
        encontrado := false
        for j := 0; j < tam2; j++{
            if vet1[i] == vet2[j] {
                encontrado = true
                break
            }
            
        }
        if !encontrado {
        contido = false
        }
    }
    if contido {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
