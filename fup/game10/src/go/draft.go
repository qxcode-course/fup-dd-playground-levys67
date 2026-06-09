package main
import "fmt"
func main() {
    n := 0
    d := 0
    a := 0
    fmt.Scan(&n, &d, &a)
    if a == d {
        fmt.Println(0)
    } else if a < d {
        fmt.Println(d - a)
    } else {
        contagem := 0
        encontrado := false
        for i := a; i < n; i++ {
           contagem++ 
           if i == d {
            encontrado = true
            break
           }
        }
        if !encontrado {
            for i := 1; i <= d; i++ {
                contagem++
            }
        }
        fmt.Println(contagem)
    }
}