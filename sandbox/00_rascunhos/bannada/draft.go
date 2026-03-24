package main
import "fmt"
func main() {
    var a float32 = 2
    var b float32 = 8
    var c float32 = 5
    fmt.Scan(&a, &b, &c)
    var z float32 = 1555
    var x float32 = 4206
    var y float32 = 761
    fmt.Scan(&z, &x, &y)
    var din float32 = 54771.87
    var gas float32 = a * z + b * x + c * y
    fmt.Scan(&din, &gas)
    fmt.Println(din - gas)
}
