package main
import "fmt"

func criarvetor (num int) []int{
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func ordenar (vetor []int) []int {
    aux := 0
    for i := 0; i < len(vetor)-1; i++{
        if vetor[i] > vetor[i+1] {
            aux = vetor[i]
            vetor[i] = vetor[i+1]
            vetor[i+1] = aux
        }
    }
    return vetor
}

func existe (vetor []int, num int) bool{
    for i := 0;i < len(vetor); i++{
        if vetor[i] == num{
            return true
        }
    }
    return false
}

func percorre (vetor []int) int {
    ordenado := 0
    menor := vetor[0]
    for i := 0; i < len(vetor); i++{
        if vetor[i] < menor {
            menor = vetor[i]
        }
    }
    ordenado = menor
    return ordenado
}

func percorrimento (vetor []int) []int {
    fila := make([]int, 0)
    for i := 0; i < len(vetor); i++ {
        fila = append(fila, percorre(vetor))
    }
    return fila
}


func main() {
    num := 0
    fmt.Scan(&num)
    vetor := criarvetor(num)
    fmt.Println(vetor)
    fila := percorrimento(vetor)
    fmt.Println(fila)
}