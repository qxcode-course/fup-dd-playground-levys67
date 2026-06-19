package main
import "fmt"
func main() {
    numero := 0
    fmt.Scan(&numero)
    if numero % 2 == 0 {
        fmt.Println("PAR")
    } else {
        fmt.Println("IMPAR")
    }
}