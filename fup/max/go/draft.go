package main
import "fmt"
func main() {
    a := 0
    b := 0
    fmt.Scan(&a,&b)
    if a == b {
        fmt.Println(a)
    } else if a > b {
        fmt.Println(a)
    } else {
        fmt.Println(b)
    }
        
    
}
