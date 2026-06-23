package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)
func main() {
    var text string
    ind := 0
    qtd := 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()
    scanner.Scan()
    indt := ""
    qtdt := ""
    indt = scanner.Text()
    scanner.Scan()
    qtdt = scanner.Text()
    ind,_ = strconv.Atoi(indt)
    qtd,_ = strconv.Atoi(qtdt)
    runas := []rune(text)
    j := 0
    for i := ind ; i < len(runas); i++ {
        if j == qtd {
            break
        }
        fmt.Printf("%c", runas[i])
        j ++
    } 
    fmt.Println("")
}