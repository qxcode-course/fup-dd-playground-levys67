package main
import "fmt"
func main() {
    a, b := 0, 0
    fmt.Scan(&a,&b)
    fmt.Printf("[ ")
    for a := a; a < b; a++ {
        fmt.Printf("%d ",a)
    }
    fmt.Printf("]\n")
}
