package main
import "fmt"
func main() {
    numero := 0
    fmt.Scan(&numero)
    cont := 0
    ehPrimo := true
    for i := 2; i <= numero; i++ {
        if numero % i == 0 {
            cont ++
        }
        if cont > 1 {
            ehPrimo = false
        }
    }
    if ehPrimo {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}