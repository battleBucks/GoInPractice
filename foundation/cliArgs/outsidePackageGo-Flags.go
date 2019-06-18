package main

import (
        "fmt"
        flags "github.com/jessevdk/go-flags"
)

var opts struct {
        Name string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
        Spanish bool `short:"s" long:"spanish" description:"A name to say hello to."`
}

func main() {
        flags.Parse(&opts) // must be used in order to get values of flag

        if opts.Spanish == true {
                fmt.Printf("Hola, %s!\n",opts.Name)
        } else {
                fmt.Printf("Hello, %s!\n",opts.Name)
        }
}
