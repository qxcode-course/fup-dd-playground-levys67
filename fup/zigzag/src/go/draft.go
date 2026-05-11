package main
import "fmt"
func main() {
    n1 := 0
    n2 := 0
    fmt.Scan(&n1, &n2)
    for i := n1; i <= n2; i++{
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("zigzag")
        } else if i % 3 == 0 && i % 5 != 0{
            fmt.Println("zig")
        } else if i % 3 != 0 && i % 5 == 0{
            fmt.Println("zag")
        } else {
            fmt.Println(i)
        }
    }
}
