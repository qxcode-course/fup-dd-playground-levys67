package main
import "fmt"

func criarvetor (qtd int) []float64 {
    vetor := make([]float64, qtd)
    for i := 0; i < len(vetor); i++ {
        fmt.Scan(&vetor[i])
    }
    return vetor
}


func calcularmedia (vetor [] float64) float64 {
    soma := 0.0
    for i := 0; i < len(vetor); i++ {
        soma += vetor[i]
    }
    media := soma / float64(len(vetor))
    return media
}

func compararmedia (media float64, vetor []float64)  []string {
    tamanhos := make([]string, 0)
    for _, v := range vetor {
        if v > media {
            tamanhos = append(tamanhos, "G")
        } else if v == media {
            tamanhos = append(tamanhos, "M")
        } else {
            tamanhos = append(tamanhos, "P")
        }
    }
    return tamanhos
}


func main() {
    qtd := 0
    fmt.Scan(&qtd)
    vetor := criarvetor(qtd)
    media := calcularmedia(vetor)
    fmt.Printf("%.2f\n",media)
    tamanhos := compararmedia(media, vetor)
    for i, v := range tamanhos {
        if i != 0 && i != len(tamanhos) {
            fmt.Print(" ")
        } 
        fmt.Print(v)
    }
    fmt.Println("")
}