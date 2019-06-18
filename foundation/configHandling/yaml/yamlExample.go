package main

import (
        "fmt"
        "github.com/kylelemons/go-gypsy/yaml"
)

func main() {
        config, err := yaml.ReadFile("conf.yaml")
        if err != nil {
                fmt.Println(err)
        }
        food, err := config.Get("food")
        quantity, err := config.GetInt("quantity")

        fmt.Printf("There are %d remaining of item: %s\n", quantity, food)
}
