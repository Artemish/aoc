package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
    "strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        phrases := strings.Split(line, " ")

        exists := map[string]bool {}

        for _, phrase := range phrases {
            if _, ok := exists[SortString(phrase)]; ok {
                total -= 1
                break
            } else {
                exists[SortString(phrase)] = true
            }
        }
        total += 1
    }
    fmt.Printf("%d\n", total)
}
