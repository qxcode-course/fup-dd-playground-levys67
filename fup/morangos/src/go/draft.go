package main
import "fmt"
func main() {
    comp1 := 0
    larg1 := 0
    comp2 := 0
    larg2 := 0
    fmt.Scan(&comp1, &larg1, &comp2, &larg2)
    area1 := comp1 * larg1
    area2 := comp2 * larg2
    if area1 > area2 {
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    }
    
}