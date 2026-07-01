package main
import "fmt"

func criarvetor (num int) []int{
    vetor := make([]int, num)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func contartazos(vetor []int) map[int]int{
    contagem := make(map[int]int)
    for  _, v := range vetor{
        contagem[v] ++
    }
    return contagem
}

func qualmaisaparece(cont map[int]int) []int{
    mais := make([]int, 0)
    val := 0
    for _, valor := range cont{
        if val < valor {
            val = valor
        }
    }    
    for chave, valor := range cont{
        if valor == val {
            mais = append(mais, chave)
        }
    }
    return mais
}

func ordenar(vetor[]int) []int {
    for i := 0; i < len(vetor); i++{
        for j := 0; j < len(vetor) - 1 - i; j++ {
            if vetor[j] > vetor[j+1] {
                vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
            }
        }
    }
    return vetor
}



func main() {
    num := 0
    fmt.Scan(&num)
    tazos := criarvetor(num)
    contagem := contartazos(tazos)
    mais := qualmaisaparece(contagem)
    mais = ordenar(mais)
    fmt.Print("[ ")
    for i := 0; i < len(mais);i++ {
        fmt.Printf("%d ", mais[i])
    }
    fmt.Println("]")
}