package main
import "fmt"
func main() {
    jog1 := 0
    jog2 := 0
    jog3 := 0
    fmt.Scan(&jog1, &jog2, &jog3)
    if jog1 == jog2 && jog2 == jog3 {
        fmt.Println("empate")
    } else if jog1 != jog2 && jog1 != jog3 {
        fmt.Println("jog1")
    } else if jog2 != jog1 && jog2 != jog3 {
        fmt.Println("jog2")
    } else {
        fmt.Println("jog3")
    }
}