
package main
import "fmt"
func main() {
    h := 0
    p := 0
    f := 0
    d := 0 // -1 para horaria e 1 para anti horaria
    fmt.Scan(&h, &p, &f, &d)
    switch d {
    case -1:
        encontrado := false
        for i := f; i >= 0; i-- {
            if i == h {
                fmt.Println("S")
                encontrado = true
                break
            } else if i == p {
                fmt.Println("N")
                encontrado = true
                break
            }

        }
        if !encontrado {
            for j := 15; j > f; j-- {
                if j == h {
                    fmt.Println("S")
                    break
                } else if j == p {
                    fmt.Println("N")
                    break
                }
            }
        }
    case 1:
        encontrado := false 
        for i := f; i <= 15; i++ {
            if i == h {
                fmt.Println("S")
                encontrado = true
                break
            } else if i == p {
                fmt.Println("N")
                encontrado = true 
                break
            } 
        }
        if !encontrado {
            for j := 0; j < f; j++ {
                if j == h {
                    fmt.Println("S")
                    encontrado = true
                    break
                } else if j == p {
                    fmt.Println("N")
                    encontrado = true
                    break
                }
            }
        }
    } 
}