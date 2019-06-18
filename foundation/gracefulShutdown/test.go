package main

import (
        "fmt"
)

func main() {
        num := []int{1,2,3,4,5,0}
        index := 0
        min := num[index]
        for i, v := range num {
                if v < min {
                        min = v
                        index = i
                }
        }
        fmt.Println(min)
        fmt.Println(index)
        num[3] = 1
        fmt.Println(num)

        sum := 0
        for _, v := range num {
                sum = sum + v
        }
        fmt.Println(sum)

        k := 4

        for i := 0; i < k; i++ {
                fmt.Println("..")
        }
}
