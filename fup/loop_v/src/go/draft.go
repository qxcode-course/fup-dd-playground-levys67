package main
import "fmt"
func main() {
    a := 0
    b := 0
    fmt.Scan(&a, &b)
    fmt.Print("[")
    for i := a; i <= b; i++{
         if i % 2 != 0 {
            fmt.Printf(" %d", i)
        }
    }
    fmt.Println(" ]")
}
