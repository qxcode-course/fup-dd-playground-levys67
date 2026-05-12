package main
import "fmt"
func main() {
    dia := 0
    hora := 0
    fmt.Scan(&dia, &hora)
    switch dia {
        case 1:{
            fmt.Println("NAO")
        }
    case 2, 3, 4, 5, 6  :{
        if (hora >= 8 && hora <= 11) || (hora >= 14 && hora <= 17){
            fmt.Println("SIM")
        } else {
            fmt.Println("NAO")
        }
    }
    case 7 :{
        if hora <= 11 && hora >= 8{
            fmt.Println("SIM")
        } else {
            fmt.Println("NAO")
        }
    }
    }
    
}
