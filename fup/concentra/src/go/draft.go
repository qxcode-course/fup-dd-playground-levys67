 package main
import "fmt"
func main() {
    var1 := 0
    var2 := 0
    fmt.Scan(&var1, &var2)
    fmt.Print("[")
    for i, j := var1, var2 ; i <= var2 && j >= var1; i, j = i + 1, j - 1{
        fmt.Printf(" %d %d", i, j)
    }
    fmt.Println(" ]")
}
