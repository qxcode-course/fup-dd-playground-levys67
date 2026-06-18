package main
import "fmt"
func main() {
    alcool := 0.0
    gasolina := 0.0
    rendimento_a := 0.0
    rendimento_g := 0.0
    fmt.Scan(&alcool, &gasolina, &rendimento_a, &rendimento_g)
    if rendimento_a > rendimento_g {
        fmt.Println("A")
    } else {
        fmt.Println("G")
    }
}