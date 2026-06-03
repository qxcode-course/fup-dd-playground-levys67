package main
import "fmt"
func main() {
    cap := 0
    ban := 0
    goi := 0
    man := 0
    fmt.Scan(&cap, &ban, &goi, &man)
    tot := ban + goi + man
    time := 0
    for i := tot; i > 0; i -= cap {
        time++
    }
    fmt.Println(time)
}