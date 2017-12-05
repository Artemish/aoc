package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        phrases := strings.Split(line, " ")

        exists := map[string]bool {}

        for _, phrase := range phrases {
            if _, ok := exists[phrase]; ok {
                total -= 1
                break
            } else {
                exists[phrase] = true
            }
        }
        total += 1
    }
    fmt.Printf("%d\n", total)
}
