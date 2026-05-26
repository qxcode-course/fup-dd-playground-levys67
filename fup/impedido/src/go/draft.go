package main
import "fmt"
func main() {
    l := 0
    r := 0
    d := 0
    fmt.Scan(&l, &r, &d)
    if r > 50 && l < r && r > d {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
