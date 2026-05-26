package main
import "fmt"
func main() {
    a := 0
    b := 0
    fmt.Scan(&a,&b)
    soma := 0
    if a <= b {
        for i := a; i <= b; i++ {
            if i % 2 == 0 {
                soma += i
            }
        }
    fmt.Println(soma)
    } else if a > b{
        fmt.Println("invalido")
    }
}
