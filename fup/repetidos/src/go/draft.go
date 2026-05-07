package main
import "fmt"
func main() {
    procurado := 0
    tamanho := 0
    fmt.Scan(&procurado, &tamanho)
    numeros := make([]int, tamanho)
    for i := 0; i < tamanho; i++ {
        fmt.Scan(&numeros[i])
    }
    achado := 0
    for _, v := range numeros {
        if procurado == v {
            achado += 1
        }
    }
    fmt.Println(achado)
}
