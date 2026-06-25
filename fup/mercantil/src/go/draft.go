package main
import "fmt"

func criarvetor(qtd int) []float64 {
    vetor := make([]float64, qtd)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func criarvetor2(qtd int) []string {
    vetor := make([]string, qtd)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func comparar3vetores(real, chute1 [] float64, chute2 []string) (int, int) {
    cont1 := 0
    cont2 := 0
    for i := 0; i < len(real) && i < len(chute1) && i < len(chute2); i++{
        if real[i] == chute1[i] || (real[i] > chute1[i] && chute2[i] == "m") || (real[i] < chute1[i] && chute2[i] == "M") {
            cont1 ++
        } else {
            cont2 ++
        }
    }
    return cont1, cont2
}

func main() {
    qtd := 0
    fmt.Scan(&qtd)
    preços := criarvetor(qtd)
    chutes1 := criarvetor(qtd)
    chutes2 := criarvetor2(qtd)
    acertos1, acertos2 := comparar3vetores(preços,chutes1,chutes2)
    if acertos1 == acertos2 {
        fmt.Println("empate")
    } else if acertos1 > acertos2 {
        fmt.Println("primeiro")
    } else {
        fmt.Println("segundo")
    }

}