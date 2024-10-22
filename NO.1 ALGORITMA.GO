package main

import "fmt"

func adalahHimpunanValid(arr []int) bool {
    if len(arr) > 100 {
        return false
    }

    set := make(map[int]bool)
    for _, num := range arr {
        if set[num] {
            return false
        }
        set[num] = true
    }
    return true
}

func main() {
    A := []int{11, 28, 33, 64, 95, 16}
    B := []int{11, 28, 33, 64, 95, 16, 100, 28, 33, 64, 95, 16}

    if adalahHimpunanValid(A) {
        fmt.Println("A adalah himpunan valid.")
    } else {
        fmt.Println("A bukan himpunan valid.")
    }

    if adalahHimpunanValid(B) {
        fmt.Println("B adalah himpunan valid.")
    } else {
        fmt.Println("B bukan himpunan valid.")
    }
}