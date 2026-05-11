package main
import "fmt"
func main() {
    n := 0
    fmt.Scan(&n)
    for i := 0; i <= n; i++{
        if i % 2 != 0{
            fmt.Println(i)
        }
    }
    for i := n; i >= 0; i-- {
        if i % 2 == 0{
            fmt.Println(i)
        }

    }
}
