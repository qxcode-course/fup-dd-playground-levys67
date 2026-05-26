package main
import "fmt"
func main() {
    val := 0
    fmt.Scan(&val)
    if val > 0 {
        fmt.Println("positivo")
    } else if val < 0 {
        fmt.Println("negativo")
    } else {
        fmt.Println("nulo")
    }
}
