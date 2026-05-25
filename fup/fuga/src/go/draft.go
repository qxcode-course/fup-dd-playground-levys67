
package main
import "fmt"
func main() {
    h := 0
    p := 0
    f := 0
    d := 0 // -1 para horaria e 1 para anti horaria
    fmt.Scan(&h, &p, &f, &d)
    if d == 1{
        for i := f; i >= 0 ; i--{
            switch i {
            case h :{
                fmt.Println("S")
                break
            }
            case p :{
                fmt.Println("N")
                break
            }

            }
        }
    }
}