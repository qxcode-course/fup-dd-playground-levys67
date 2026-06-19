package main
import "fmt"
func main() {
    numero1 := 0
    numero2 := 0
    fmt.Scan(&numero1, &numero2)
    if (numero1 % 3 == 0 && numero2 % 3 == 0) || (numero1 % 5 == 0 && numero2 % 5 == 0) {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}