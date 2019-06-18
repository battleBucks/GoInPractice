package main

import (
        "fmt"
)

func main() {
        arr := [][]int32{
                {1,1,1,0,0,0},
                {0,1,0,0,0,0},
                {1,1,1,0,0,0},
                {0,0,2,4,4,0},
                {0,0,0,2,0,0},
                {0,0,1,2,4,0},
        }
        arr2 := [][]int32{
                {-9,-9,-9,1,1,1},
                {0,-9,0,4,3,2},
                {-9,-9,-9,1,2,3},
                {0,0,8,6,6,0},
                {0,0,0,-2,0,0},
                {0,0,1,2,4,0},
        }

        _ = hourglassSum(arr)
        _ = hourglassSum(arr2)

}

func hourglassSum(arr [][]int32) int32 {
        max := int32(-999999)
        min := int32(999999)

        for i := 0; i < (len(arr) - 2); i++ {
                for j := 0; j < (len(arr[0]) -2); j++ {
                        sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
                        if max < sum {
                                max = sum
                        }
                        if min > sum {
                                min = sum
                        }
                        fmt.Println("Sum is: ",sum)
                }
        }
        fmt.Printf("Max: %d\t|\tMin: %d\n", max, min)
        return max
}
