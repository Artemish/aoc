package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var nums []int

    for scanner.Scan() {
        n, err := strconv.Atoi(scanner.Text())
        if err == nil {
            nums = append(nums, n)
        }
    }

    steps := 0

    for pos := 0; pos >= 0 && pos < len(nums); {
        steps += 1
        nums[pos] += 1
        pos += nums[pos] - 1
    }

    fmt.Printf("%d\n", steps)
}
