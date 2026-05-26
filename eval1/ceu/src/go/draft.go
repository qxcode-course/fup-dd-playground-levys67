 package main
import "fmt"
func main() {
    pedra := 0
    fmt.Scan(&pedra)
    fmt.Print("[")
    for i := 0; i <= 10; i++ {
        if i != pedra && i != 10 {
            fmt.Printf(" %d", i)
        }
        if i == 10 && i != pedra{
            fmt.Print(" ceu")
        }
    }    
    fmt.Println(" ]")
}
