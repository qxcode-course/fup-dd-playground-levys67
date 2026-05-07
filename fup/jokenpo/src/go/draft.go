 package main
import "fmt"
func main() {
    jog1 := ""
    jog2 := ""
    fmt.Scan(&jog1, &jog2)
    if jog1 == jog2 {
        fmt.Println("empate")
        // R ganha de S que ganha de P que ganha de R
    } else if (jog1 == "R" && jog2 == "S") || (jog1 == "S" && jog2 == "P") || (jog1 == "P" && jog2 == "R") {
        fmt.Println("jog1")
    } else {
        fmt.Println("jog2")
    }
}
