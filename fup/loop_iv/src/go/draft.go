package main
import "fmt"
func main() {
    a := 0
    b := 0
    fmt.Scan(&a, &b)
    fmt.Print("[")
    if a > b {
        for i := a; i  > b; i-- {
            fmt.Printf(" %d",i)
        }
    } else if a < b {
        for i := a; i < b; i++ {
                fmt.Printf(" %d",i)
        }
    }
    fmt.Println(" ]") 
}