package main
import "fmt"
func main() {
    n := 0
  fmt.Scan(&n)
  lista := make([]int,n)
  for i := 0; i < n; i++{
    fmt.Scan(&lista[i])
  }
if n == 0 {
    fmt.Print("\n")
}

  for _, valor := range lista{
    fmt.Println(valor)
  } 
    


}
