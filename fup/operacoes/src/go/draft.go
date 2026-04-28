package main
import "fmt"
func main() {
    a := 0
    b := 0
    fmt.Scan(&a, &b)
    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    div :=float64(a) / float64(b)
    fmt.Printf("%.2f\n",div)
    fmt.Println(a % b)
}
