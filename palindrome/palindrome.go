package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    words := []string{}

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    makePalindromes(words)
    fmt.Println(words)
}

func makePalindromes(words []string) {
    for i, w := range words {
        words[i] = makePalindrome(w)
    }
}

func makePalindrome(s string) string {
    counts := make(map[rune]int)
    result := make([]rune, len(s))
    odd := 0
    var oddRune rune

    // count letters
    for _, r := range s {
        counts[r]++
    }

    // populate result rune slice
    left, right := 0, len(s)-1
    for k, v := range counts {
        if v % 2 == 1 {
            odd++
            oddRune = k
        }
        for v > 1 {
            result[left], result[right] = k, k
            left++
            right--
            v -= 2
        }
    }

    if odd > 1 {
        return "-1"
    }

    if odd == 1 {
        result[left] = oddRune
    }

    return string(result)
}

