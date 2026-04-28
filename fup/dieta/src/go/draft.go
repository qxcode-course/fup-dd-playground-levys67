package main
import "fmt"
func main() {
    dias := 0
    fmt.Scan(&dias)
    con := make([]int, dias)
    for i := 0; i < dias; i++ {
        fmt.Scan(&con[i])
    }
    soma := 0
    for _, v := range con {
        soma += v
    }
    media := soma / dias
    fmt.Printf("%.1f\n", float64(media))
}
