package main
import "fmt"
func main() {
    a := 0
    fmt.Scan(&a)
    if a >= 0 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
