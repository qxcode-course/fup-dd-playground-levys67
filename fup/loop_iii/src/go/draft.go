package main

import "fmt"

func main() {
	n := 0
	m := 0
	fmt.Scan(&n, &m)
    fmt.Print("[")
	for i := n; i > m; i-- {
		fmt.Printf(" %d",i)
	}
    fmt.Println(" ]")
}
