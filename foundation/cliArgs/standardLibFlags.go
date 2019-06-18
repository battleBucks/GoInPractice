package main

import (
        "fmt"
        "flag"
)

var name = flag.String("name", "World", "a name to say hello to")
var greet = "Hello"

var Italian bool

func init() {
        flag.BoolVar(&Italian, "italian", false, "Use Italian language")
        flag.BoolVar(&Italian, "i", false, "Use Italian language")
}

func main() {
        flag.Parse() // must be used in order to get values of flag

        /*
        if Italian == true {
                fmt.Printf("Ciao, %s!\n",*name)
        } else {
                fmt.Printf("Hello, %s!\n",*name)
        }
        */

        if Italian {
                greet = "Ciao"
        }
        fmt.Printf("%s, %s!\n", greet, *name)
}
