package main
import "fmt"
func main() {
    var a int = 12
    var b int = 6
    fmt.Scan(&a, &b)
    var div int = a / b
    fmt.Print(div)
    var res int = a % b
    fmt.Print(" ")
    fmt.Println(res)
}
