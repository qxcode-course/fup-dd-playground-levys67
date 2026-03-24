 package main
import "fmt"
func main() {
    var a float64 
    var b float64 
    var c float64 
    var z float64 
    var x float64 
    var y float64 
    fmt.Scan(&a, &b, &c)
    fmt.Scan(&z, &x, &y) 
    var din float64  
    fmt.Scan(&din)
    var gas float64 = (a * z) + (b * x) + (c * y)
    fmt.Printf("%.2f\n",din - gas)
}
