package main
import "fmt"
func main() {
    n := 0
	fmt.Scan(&n)
	lista := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lista[i])
	}
	fmt.Print("[")
	for i, v := range lista {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}
