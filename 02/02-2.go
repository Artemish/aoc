package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        tokens := strings.Split(line, " ")

        l := len(tokens)

        var n int
        var err error
        var numbers = make([]int, l)

        for index, num := range tokens {
            if n, err = strconv.Atoi(num); err == nil {
                numbers[index] = n
            }
        }

        for i := 0; i < l; i++ {
            for j := (i+1); j < l; j++ {
                n_i, n_j := numbers[i], numbers[j]
                if (n_i % n_j) == 0 {
                    total += n_i / n_j
                    break
                } else if (n_j % n_i) == 0 {
                    total += n_j / n_i
                    break
                }
            }
        }
    }
    fmt.Printf("%d\n", total)
}
