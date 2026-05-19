
package main
import "fmt"
func main() {
    h := 0
    p := 0
    f := 0
    d := 0 // -1 para horaria e 1 para anti horaria
    fmt.Scan(&h, &p, &f, &d)
    if d < 0 && f < p {
        for i := f; i >= 0; i--{
            if i == h {
                fmt.Println("S")
                break
            } else if i == p {
                fmt.Println("N")
                break
            }
            
        } 
       
    }
}