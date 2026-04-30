package main
import "fmt"
func main() {
    a := 0
    b := 0
    fmt.Scan(&a, &b)
    if a < b {
        soma := 0
        for i := a; i <= b; i++ {
            if i % 2 == 0 {
                soma += i
            }
        }
        fmt.Println(soma)
    } else if a > b {
        fmt.Println("invalido")
    } else if a == b {
        fmt.Println(a)
    }
}
