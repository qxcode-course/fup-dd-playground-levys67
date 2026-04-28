package main
import "fmt"
func main() {
    lista := make([]int, 5)
    for i := 0; i < 5; i++ {
        fmt.Scan(&lista[i])
    }
    menor := lista[0]
    for _, v := range lista {
        if v < menor {
            menor = v
        }
    }
    fmt.Println(menor)
}
