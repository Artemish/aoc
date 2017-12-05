package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    l := len(text) - 1
    total := 0

    for i := 0; i < l; i++ {
        next := (i + (l / 2))
        if text[i] == text[next % l] {
            total += int(text[i] - byte('0'))
        }
    }

    fmt.Println(total)
}
