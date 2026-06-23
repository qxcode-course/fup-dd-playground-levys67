package main
import "fmt"

func criarvetor(tam int) []int {
    vetor := make([]int, tam)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }
    return vetor
}

func separarpares(vetor []int) ([]int, []int) {
    machos := make([]int, 0)
    femeas := make([]int, 0)
    for _, v := range vetor{
        if v > 0 {
            machos = append(machos, v)
        } else {
            femeas = append(femeas, v)
        }
    }
    return machos, femeas
}

func contarpares (machos,femeas []int) int {
    contagem := 0
    for _, v := range machos{
        for _, k := range femeas{
            if v * (-1) == k {
                contagem ++
            }
        }
    }
    return contagem
}


func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    //fmt.Println(vetor)
    machos, femeas := separarpares(vetor)
    fmt.Println(machos, femeas)
    contagem := contarpares(machos, femeas)
    fmt.Println(contagem)

}