package main

import (
    "fmt"
)

func answer(target int) int {
    MAX_SIZE := 256
    i, j := 128,128
    n := 1

    arr := make([][]int, MAX_SIZE)
    for index := range arr {
        arr[index] = make([]int, MAX_SIZE)
    }
    arr[i][j] = 1

    for {
        i += 1
        n += 2

        arr[i][j] = (arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
                     arr[i][j-1] + arr[i][j+1] +
                     arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1])

        if arr[i][j] > target { return arr[i][j] }

        var x int

        for x = 0; x < n-2; x++ {
            j += 1

            arr[i][j] = (arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
                         arr[i][j-1] + arr[i][j+1] +
                         arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1])

            if arr[i][j] > target { return arr[i][j] }
        }

        for x = 0; x < n - 1; x++ {
            i -= 1
            arr[i][j] = (arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
                         arr[i][j-1] + arr[i][j+1] +
                         arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1])

            if arr[i][j] > target { return arr[i][j] }
        }

        for x = 0; x < n-1; x++ {
            j -= 1
            arr[i][j] = (arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
                         arr[i][j-1] + arr[i][j+1] +
                         arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1])

            if arr[i][j] > target { return arr[i][j] }
        }

        for x = 0; x < n-1; x++ {
            i += 1
            arr[i][j] = (arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
                         arr[i][j-1] + arr[i][j+1] +
                         arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1])

            if arr[i][j] > target { return arr[i][j] }
        }
    }
}

func main() {
    var target int
    fmt.Scan(&target)
    fmt.Printf("%d\n", answer(target))
}
