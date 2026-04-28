package main
import "fmt"
func main() {
    n := 0
    d := 0
    fmt.Scan(&n, &d)
    div := n / d
    rest := n % d
    decimal := float64(n) / float64(d)
    fmt.Printf("%d\n%d\n%.2f\n", div, rest, decimal)
}
