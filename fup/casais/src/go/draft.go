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

func contarpares (animais[]int ) int {
    cont := 0
    mapa1 := make(map[int]int)
    mapa2 := make(map[int]int)
    for i := 0; i < len(animais);i++{
        if animais[i] > 0{
            mapa1[i] ++
        } else {
            mapa2[i] ++
        }
    }
    for chave, valor := range mapa1 {
        esp := false 
        for key := range mapa2 {
            if key * -1 == chave {
                esp = false 
            }

        }
        if esp && valor >  0 && mapa2[chave *-1] > 0 {
            cont++
            valor --
            mapa2[chave * -1] --
        }
    }
    return cont
}

func main() {
    tam := 0
    fmt.Scan(&tam)
    vetor := criarvetor(tam)
    fmt.Println(vetor)
    contagem := contarpares(vetor)
    fmt.Println(contagem)
}