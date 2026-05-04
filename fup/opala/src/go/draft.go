package main
import "fmt"
func main() {
    vel := 0
    time := 0
    cons := 0.0
    fmt.Scan(&vel, &time, &cons)
    t_hora := float64(time) / 60
    dist := float64(vel) * t_hora
    desm := dist / float64(cons)
    fmt.Printf("%.2f\n", float64(desm))
}
