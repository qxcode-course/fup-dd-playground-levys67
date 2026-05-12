package main
import "fmt"
func main() {
    figdoalbum := 0
    figtotal := 0
    inteiros := 0
    fmt.Scan(&figdoalbum, &figtotal)
    lista := make([]int, figtotal)
    for i := 0; i < figtotal; i++{
        fmt.Scan(inteiros)
    }
    fmt.Print(lista)
}
