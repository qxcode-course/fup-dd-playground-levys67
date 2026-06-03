package main
import "fmt"

func criarepreenchervetor(n int) []int {
    lista := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&lista[i])
    }
    return lista
}

func contarrepetidos(lista []int) []int {
    var repetidos []int
    for i := 0; i < len(lista)-1; i++ {
        if lista[i] == lista[i+1] {
            repetidos = append(repetidos, lista[i])
        }
    }
    return repetidos

}

func main() {
    figdoalbum := 0
    figtotal := 0
    //inteiros := 0
    fmt.Scan(&figdoalbum, &figtotal)
    lista := criarepreenchervetor(figtotal)
    //fmt.Print(lista)
    repetidos := contarrepetidos(lista)
    fmt.Println(repetidos)
}
