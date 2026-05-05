package main
import  "fmt"
import  "bufio"
import  "os"

func main() {
    var texto string
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto = scanner.Text()
    // fmt.Scanln(&texto)
    listdetexto := []rune(texto)
    for i := len(texto)-1; i >= 0; i-- {
        fmt.Printf("%c", listdetexto[i])
    }
    fmt.Println()

}
