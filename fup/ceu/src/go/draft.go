package main
import "fmt"
func main() {
    num := 0
    fmt.Scan(&num)
    fmt.Print("[")
    for i := 0; i <= 10; i++ {
        if i != num && i != 10 {
            fmt.Printf(" %d", i)
        }
    }
    for i := 0; i <= 10; i++ {
        if i != num && i == 10{
            fmt.Print(" ceu")
        }
    }
    fmt.Println(" ]")
}
