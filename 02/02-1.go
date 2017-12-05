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
        numbers := strings.Split(line, " ")

        var i, min, max int
        var err error

        if i, err = strconv.Atoi(numbers[0]); err == nil {
            min = i
            max = min
        }

        for _, num := range numbers {
            if i, err = strconv.Atoi(num); err == nil {
                if i < min {
                    min = i
                }

                if i > max {
                    max = i
                }
            }
        }
    }
    fmt.Printf("%d\n", total)
}
