package main
import "fmt"
func main() {
    val := 0
    par := 0
    fmt.Scan(&val, &par)
    cdpar := float64((val / par))
    cdjur := cdpar * 0.05
    cdmom := cdpar + (cdjur * (float64(par) - 1))
    if par == 1 {
        fmt.Printf("%.2f\n", float64(val))
    }else if par != 1 {
        fmt.Printf("%.2f\n", cdmom)
    }
    total := cdmom * float64(par)
    if par == 1 {
        fmt.Printf("%.2f\n", float64(val))
    } else if par != 1 {
        fmt.Printf("%.2f\n", total)
    }

}
