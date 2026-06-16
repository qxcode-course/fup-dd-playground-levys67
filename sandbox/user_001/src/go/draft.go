package main
import "fmt"
func main() {
	nomes := make(map[string]int)

	for {
		palavra := ""
		fmt.Scan(&palavra)
		if palavra == "pare"{
			break
		}
		nomes[palavra] += 1
	}
fmt.Println(nomes)	
}