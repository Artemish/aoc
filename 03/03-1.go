package main

import (
    "fmt"
    "math"
)

func main() {
    var i int
    fmt.Scan(&i)

    n := 1

    for ; i >= n*n; n += 2 { }

    fmt.Printf("n = %d\n", n)

    var dist, pivot int

    dist = (n - 1) / 2

    if (i > n*n - n + 1) { // Bottom
        pivot = n*n - (n-1)/2
        fmt.Printf("A Dist(%d), Pivot(%d)\n", dist, pivot)
        dist += int(math.Abs(float64(i - pivot)))
    } else if (i > n*n - 2*n + 2) { // Left
        pivot = n*n - (3*(n-1))/2
        fmt.Printf("B Dist(%d), Pivot(%d)\n", dist, pivot)
        dist += int(math.Abs(float64(i - pivot)))
    } else if (i > n*n - 3*n + 3) { // Top
        pivot = n*n - (5*(n-1))/2
        fmt.Printf("C Dist(%d), Pivot(%d)\n", dist, pivot)
        dist += int(math.Abs(float64(i - pivot)))
    } else { // Right
        pivot = n*n - (7*(n-1))/2
        fmt.Printf("D Dist(%d), Pivot(%d)\n", dist, pivot)
        dist += int(math.Abs(float64(i - pivot)))
    }

    fmt.Printf("Distance: %d\n", dist)
}
