package main

import "fmt"
func main() {
    wifi := 0
    login := 0
    admin := 0
    fmt.Scan(&wifi,&login,&admin)
    if wifi == 0 {
        fmt.Println("you must connect to wifi")
    } else if login == 0 {
        fmt.Println("you need to login first")
    } else if admin == 0{
        fmt.Println("you must to login as admin")
    } else {
        fmt.Println("done")
    }
}

