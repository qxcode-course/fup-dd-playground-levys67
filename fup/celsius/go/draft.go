package main
import "fmt"
func main() {
    var cel float64 = 0
    fmt.Scan(&cel)
    var fah float64 = (cel * 1.8) + 32
    fmt.Printf("%.6f\n",fah)
}
