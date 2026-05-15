 package main
import "fmt"
func main() {
    var1 := 0
    var2 := 0
    fmt.Scan(&var1, &var2)
    fmt.Print("[")
    for i := var1 ; i <= var2 * 2 ; i++{
        fmt.Print(i)
    }
}
