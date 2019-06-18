package main

import (
        "fmt"
        "os"
        "github.com/urfave/cli"
)

func main() {
        app := cli.NewApp()
        app.Name = "getAFewOptions"
        app.Usage = "Greet User and suck in butt farts"
        app.Flags = []cli.Flag{
                cli.StringFlag {
                        Name: "name, n",
                        Value: "World",
                        Usage: "Who to say hello to.",
                },cli.IntFlag{
                        Name: "age, a",
                        Value: 0,
                        Usage: "make fun of age",
                },cli.BoolFlag{
                        Name: "suck, s",
                        Usage: "sucking in butt farts",
                },
        }
        app.Action = func(c *cli.Context) error {
                name := c.GlobalString("name")
                suck := c.GlobalBool("suck")
                age := c.GlobalInt("age")
                greet := ""
                if suck {
                        fmt.Println("SUCK IN MY BUTTFART!!!")
                } else {
                        if age > 0 {
                                greet = fmt.Sprintf("Hello %s, who is %d years old!\n", name, age)
                        } else {
                                greet = fmt.Sprintf("Hello %s!\n", name)
                        }
                        fmt.Printf(greet)
                }
                return nil
        }

        app.Run(os.Args)
}
