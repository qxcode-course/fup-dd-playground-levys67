package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    var text string
    ind := 0
    qtd := 0
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()
    fmt.Sscan(text, &ind, &qtd)
    runas := []rune{text}
    fmt.Println(runas)
}