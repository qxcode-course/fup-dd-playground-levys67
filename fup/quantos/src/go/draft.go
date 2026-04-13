package main
import "fmt"
func main() {
    jog1 := 0
    jog2 := 0
    jog3 := 0
    fmt.Scan(&jog1,&jog2,&jog3)
    if jog1 == jog2 && jog1 == jog3{
        fmt.Println(3)
    } else if jog1 == jog2 || jog2 == jog3 || jog1 == jog3 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}
