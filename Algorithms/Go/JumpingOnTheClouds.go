package main

import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    clouds := make([]int, n)
    for i := range clouds {
        fmt.Scanf("%d", &clouds[i])
    }
    moves := 0;
    i := 0
    for i < n-1 {
        moves++
        if i + 2 < n && clouds[i + 2] == 0 {
            i += 2
        } else {
            i++
        }
    }
    fmt.Println(moves)
}
