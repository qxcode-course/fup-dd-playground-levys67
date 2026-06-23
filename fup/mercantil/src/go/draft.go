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

func compararvetores1(vetor1, vetor2 [] float64) []string {
    res := make([]string, 0)
    for i := 0; i < len(vetor1) && i < len(vetor2); i++{
        if vetor1[i] == vetor2[i] {
            res = append(res, "A")
        } else if vetor1[i] > vetor2[i] {
            res = append(res, "M")
        } else {
            res = append(res, "m")
        }
    }
    return res
}

func quemganhou (result, chute []string) (cont1,cont2 int){
    cont1 := 0
    cont2 := 0
    for i := 0; i < len(result) && i < len(chute); i++ {
        if 
    } 
    
}


func main() {
    qtd := 0
    fmt.Scan(&qtd)
    preços := criarvetor(qtd)
    chutes1 := criarvetor(qtd)
    chutes2 := criarvetor2(qtd)
    result := compararvetores1(preços,chutes1)
    fmt.Print(result,chutes2)

}