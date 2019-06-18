package main

import (
        "fmt"
        "gopkg.in/gcfg.v1"
)

type Person struct {
        First string
        Last string
        Age int
}

func main() {
        config := struct {
                Person
        }{}

        err := gcfg.ReadFileInto(&config, "joe.ini")
        if err != nil {
                fmt.Println("Failed to parse config file %s", err)
        }
        fmt.Println(config.Person.First)
        fmt.Println(config.Person.Last)
        fmt.Println(config.Person.Age)
}
