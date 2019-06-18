package main

import (
        "encoding/json"
        "fmt"
        "os"
)

type Grayson struct {
        First string
        Last string
        Age int
}

func main() {
        f, _ := os.Open("conf.json")
        defer f.Close()

        decoder := json.NewDecoder(f)
        gray := Grayson{}
        err := decoder.Decode(&gray)
        if err != nil {
                fmt.Println("error", err)
        }
        fmt.Printf("Grayson has a name: %s %s and a %d inch nose\n", gray.First, gray.Last, gray.Age)
}

