package main

import (
    "fmt"
)

func NextPermutation(s []int) []int {
    if len(s) <= 1 {
        return s
    }

    var pivot int
    sorted := []int{s[len(s)-1]}

    for pivot = len(s)-2; pivot > 0; pivot-- {
        if s[pivot] < s[pivot+1] {
            break
        }

        sorted = append(sorted, s[pivot])
    }

    tmp := s[pivot]
    result := append(append(append(s[:pivot], s[pivot+1]), tmp), sorted[:len(sorted)-1]...)
    // s[:pivot] : s[pivot+1] : bubbledown(s[pivot], sorted)

    for pivot = pivot+1 ; pivot < len(result)-1; pivot++ {
        if result[pivot] <= result[pivot+1] {
            break
        }

        result[pivot], result[pivot+1] = result[pivot+1], result[pivot]
    }

    return result
}

func main() {
    sequence := []int{1,2,5,4,3,2,1}
    fmt.Println(NextPermutation(sequence))
    next := sequence

    for i:=0; i<1000000; i++ {
        next = NextPermutation(next)
        //fmt.Println(next)
    }
}
