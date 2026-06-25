  package main
import "fmt"



func main() {
  char := ""
  valor := 0
  for {
    fmt.Scanf("%d%c",&valor, &char)
    fmt.Printf("%d", valor)
    if char == "\n"{
      break
    }

  }

}