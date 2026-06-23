package main
import "fmt"
func main() {
    N := 0
    V := ""
    fmt.Scan(&N)

    for i := 0; i < N ; i++{

        fmt.Scan(&V)
        runas := []rune (V)
        soma := 0
        for j := 0; j < len(runas); j++{

            switch runas[j] {
            case '1' :
                soma += 2
            case '2' :
                soma += 5
            case '3' :
                soma += 5
            case '4' :
                soma += 4
            case '5' :
                soma += 5
            case '6' :
                soma += 6
            case '7' :
                soma += 3
            case '8' :
                soma += 7
            case '9' :
                soma += 6
            case '0' :
                soma += 6
            }
        }
        fmt.Printf("%d leds\n",soma )
    }

}