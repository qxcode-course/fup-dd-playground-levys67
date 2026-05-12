package main
import "fmt"
func main() {
    dia := 0
    hora := 0
    fmt.Scan(&dia, &hora)
    if dia == 1 {
        fmt.Println("NAO")
    } else if (dia > 1 && dia < 7) && (hora >= 8 && hora <= 11) || (hora >= 14 && hora <= 17) {
        fmt.Println("SIM")
    } else if dia == 7 && (hora >= 8 && hora <= 11) {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
