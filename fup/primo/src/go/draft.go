package main
import "fmt"
func main() {
    numero := 0
    fmt.Scan(&numero)
    for i := 1; i <= numero; i++ {
        if numero % i == 0 && numero == i {
            fmt.Println(1)
        }
    }
}