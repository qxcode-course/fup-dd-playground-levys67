package main
import "fmt"
func main() {
    profundidade := 0
    salto := 0
    escorrego := 0
    fmt.Scan(&profundidade, &salto, &escorrego)
    altura := 0
    for {
        fmt.Printf("%d ", altura)
        altura += salto
        if altura >= profundidade {
            fmt.Println("saiu")
            break
        }
        if altura < 0 {
            fmt.Println("morreu")
            break
        }
        fmt.Println(altura)
        altura -= escorrego
        salto -= 10
    }
}
