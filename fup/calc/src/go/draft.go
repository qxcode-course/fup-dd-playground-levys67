package main
import "fmt"
func main() {
    a := 0
    b := 0
    var op string = ""
    fmt.Scan(&a,&b,&op)
    if op == "+" {
       fmt.Println(a + b) 
    } else if op == "-"{
        fmt.Println(a - b)
    } else if op == "*" {
        fmt.Println(a * b)
    } else if op == "/"{
        fmt.Println(a / b)
    }
}
