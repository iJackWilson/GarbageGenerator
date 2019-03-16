package main

import ("fmt"
	"log"
	"os"
	"strings"
	"flag")

func main() {
    filenamePtr := flag.String("filename", "output.txt", "Name of output file")
    lengthPtr := flag.Int("length", 100, "Length of garbage")

    flag.Parse()

    fmt.Println("Creating", *lengthPtr, "characters of garbage and outputting to", *filenamePtr)

    a2 := strings.Repeat("A", *lengthPtr)

    file, err := os.Create(*filenamePtr)
    if err != nil {
	    log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    fmt.Fprintf(file, a2)
}
