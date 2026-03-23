package main
import "fmt"
func main() {
    var a int = 2
    var b int = 3
    fmt.Scan(&a, &b)
    var div int = a / b 
    var res int = a - b
    fmt.Print(div, res)
}
