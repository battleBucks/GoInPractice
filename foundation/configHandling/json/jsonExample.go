package main

import (
        "encoding/json"
        "fmt"
        "os"
)

type Parser struct {
        First string
        Last string
        Age int
}

func main() {
        f, _ := os.Open("conf.json")
        defer f.Close()

        decoder := json.NewDecoder(f)
        parsed := Parser{}
        err := decoder.Decode(&parsed)
        if err != nil {
                fmt.Println("error:",err)
        }
        fmt.Println("Config File data--------------------")
        fmt.Printf("First name: %s\tLast name: %s\tAge: %d\n", parsed.First, parsed.Last, parsed.Age)
}
