/* Given a tree, return the number of paths that sum to a target
 */

/* Input
[1,0,1,1,2,0,-1,0,1,-1,0,-1,0,1,0]
2
*/

package main

import (
    "fmt"
)

type Node struct {
    Val int
    Left, Right *Node
}

func main() {
    tree := makeTree([]int{0,1,0,1,1,2,0,-1,0,1,-1,0,-1,0,1,0}, 1)
    fmt.Println(sumPaths(tree, 2))
}

func makeTree(arr []int, i int) *Node {
    if i >= len(arr) { return nil }
    return &Node{Val: arr[i], Left: makeTree(arr, 2*i), Right: makeTree(arr, 2*i+1)}
}

func sumPaths(root *Node, sum int) int {
    count := 0

    count += sumPathsFrom(root, sum)
    if root.Left != nil {count += sumPaths(root.Left, sum)}
    if root.Right != nil {count += sumPaths(root.Right, sum)}

    return count
}

func sumPathsFrom(u *Node, sum int) int {
    if u == nil { return 0 }

    result := 0

    if sum-u.Val == 0 { result++ }
    result += sumPathsFrom(u.Left, sum-u.Val)
    result += sumPathsFrom(u.Right, sum-u.Val)

    return result
}
