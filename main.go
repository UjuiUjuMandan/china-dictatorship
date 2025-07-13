// main.go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    data, err := os.ReadFile("README.adoc")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(string(data))
}
