package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(2.0)
    s := float64(0.0)
    
    for {
        z = z - (z*z - x)/(2*z)
        if math.Abs(s-z) < 1e-15 {
            break
        }
        s = z
    }
    
    return s
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}