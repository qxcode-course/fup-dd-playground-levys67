package main
import "fmt"
func main() {
    val1 := 0
    val2 := 0
    val3 := 0
    fmt.Scan(&val1,&val2, &val3)
    if (val1 > val2 && val1 < val3) || (val1 > val3 && val1 < val2) {
        fmt.Println(val1)
    } else if (val2> val1 && val2 < val3) || (val2 > val3 && val2< val1) {
        fmt.Println(val2)
    } else if (val3 > val1 && val3 < val2) || (val3 > val2 && val3 < val1) {
        fmt.Println(val3)
    }
}
