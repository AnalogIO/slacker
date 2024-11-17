package main

import (
	"fmt"
)


func main() {
    user := getUserByEmail("maod@cafeanalog.dk")
    fmt.Println("User: ", user)
}
