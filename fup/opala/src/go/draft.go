package main
import "fmt"
func main() {
    velm := 0.0
    time := 0.0
    cons := 0.0
    fmt.Scan(&velm,&time,&cons)
    timeh := time / 60
    traj := velm * timeh
    desem := traj / cons
    fmt.Printf("%.2f\n", float64(desem))
}
