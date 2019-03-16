package main

import ("fmt"
	"log"
	"os"
	"strings")

func main() {
    var i int
    fmt.Println("Enter length of garbage: ")
    fmt.Scan(&i)
    fmt.Println("Creating", i, "characters of garbage")
    a2 := strings.Repeat("A", i)
    //fmt.Println(a2)
    file, err := os.Create("result.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    fmt.Fprintf(file, a2)

}
