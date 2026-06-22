package main
import "fmt"

func criarvetor (tam int) []int{
    vetor := make([]int, tam)
    for i := 0; i < tam; i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func contatrebeldes(vetor []int) (int,int){
    soldados := 0
    rebeldes := 0
    for _, v := range vetor{
        if v % 2 == 0 {
            rebeldes += v
        } else {
            soldados += v
        }
    }
    return soldados, rebeldes
}

func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    soldados ,rebeldes := contatrebeldes(vetor)
    if soldados == rebeldes {
        fmt.Println("empate")
    } else if soldados > rebeldes {
        fmt.Println("soldados")
    } else {
        fmt.Println("rebeldes")
    }
}