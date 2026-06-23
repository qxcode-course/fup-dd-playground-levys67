 package main
import "fmt"

func criarvetor (tam int) []int{
    vetor := make([]int, tam)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func separarnumeros (vetor []int) ([]int, []int) {
    impares := make([]int, 0)
    pares := make([]int, 0)
    for _, v := range vetor {
        if v % 2 == 0{
            pares = append(pares, v)
        } else {
            impares = append(impares, v)
        }
    }
    return impares, pares
}


func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    impares, pares := separarnumeros(vetor)
    //fmt.Println(impares, pares)
    fmt.Print("[ ")
    for i := 0; i < len(impares); i++{
        fmt.Printf("%d ", impares[i])
    }
    fmt.Println("]")
    fmt.Print("[ ")
    for i := 0; i < len(pares); i++{ 
        fmt.Printf("%d ",pares[i])
    }
    fmt.Println("]")
}