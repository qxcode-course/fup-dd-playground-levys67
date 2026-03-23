package main
import "fmt"
func main() {
    var a float32 = 12
    var b float32 = 6
    fmt.Scan(&a, &b)
    var c float32 = ((a + b) / 2)
    fmt.Printf("%.1f\n", c)
}
