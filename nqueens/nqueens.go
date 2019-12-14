//The n-queens puzzle is the problem of placing n queens on an n×n chessboard such
//that no two queens attack each other.
//
//Given an integer n, return all distinct solutions to the n-queens puzzle.
//
//Each solution contains a distinct board configuration of the n-queens'
//placement, where 'Q' and '.' both indicate a queen and an empty space
//respectively.

package main

import (
    "fmt"
)

func main() {
    n := 4
    board := make([][]string, n)
    cols := make(map[int]bool)
    diagU := make(map[int]bool)
    diagD := make(map[int]bool)

    for i:=0; i<n; i++ {
        board[i] = make([]string, n)
        for j:=0; j<n; j++ {
            board[i][j] = "."
        }
    }

    nqueens(board, n, 0, cols, diagU, diagD)
}

func nqueens(board [][]string, n, i int, cols, diagU, diagD map[int]bool) {
    if i == n {
        for _, v := range board {
            fmt.Println(v)
        }
        fmt.Println()
        return
    }

    for j:=0; j<n; j++ {
        if !cols[j] && !diagU[j-i] && !diagD[j+i] {
            cols[j] = true
            diagU[j-i] = true
            diagD[j+i] = true
            board[i][j] = "Q"
            nqueens(board, n, i+1, cols, diagU, diagD)
            board[i][j] = "."
            cols[j] = false
            diagU[j-i] = false
            diagD[j+i] = false
        }
    }
}
