// main.go
package main

import (
    _ "embed"
    "fmt"
    "log"
)

//go:embed README.adoc
var data string

func main() {
    if data == "" {
        log.Fatal("README.adoc content is empty")
    }
    fmt.Print(data)
}
