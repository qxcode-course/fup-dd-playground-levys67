package main
import "fmt"

func criarvetor(qtd int) []int{
    vetor := make([]int, qtd)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func contarsoma(vetor []int) int {
    soma := 0
    contarAS := 0
    for i := 0; i < len(vetor); i++{
        if vetor[i] == 1 {
            contarAS ++
        }
        if  vetor[i] == 11 || vetor[i] == 12 || vetor[i] == 13{
            soma += 10
        } else if vetor[i] == 1{
            soma += 11
        } else {
            soma += vetor[i]
        }
        
    }
    if soma > 21 && contarAS <= 2{
        soma -= 10*contarAS
    } else if soma > 21 && contarAS > 2 {
        soma -= 10*(contarAS -1) 
    }
    return soma
}

func main() {
    qtd := 0
    fmt.Scan(&qtd)
    cartas := criarvetor(qtd)
    somatoria := contarsoma(cartas)
    fmt.Println(somatoria)
}