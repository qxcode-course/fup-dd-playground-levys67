package main

import "fmt"

func main() {
	var1 := 0
	var2 := 0
	var3 := 0
	fmt.Scan(&var1, &var2, &var3)
	soma1 := var2 + var3
	soma2 := var1 + var3
	soma3 := var1 + var2
	if var1 >= soma1 || var2 >= soma2 || var3 >= soma3 {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}
