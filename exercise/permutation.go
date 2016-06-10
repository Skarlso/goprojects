package main

import (
    "strings"

    "github.com/skarlso/goutils/arrayutils"
)

var locations = []string{"1", "2", "3", "4"}
var permutations = make([]string, 0)

func permute(s []string) {
    for i := 0; i < fact(len(s)); i++ {
        innerI := i % len(s)
        for j := 0; j < len(s); j++ {
            s[innerI], s[j] = s[j], s[innerI]
            if !arrayutils.ContainsString(permutations, strings.Join(s, " ")) {
                permutations = append(permutations, strings.Join(s, " "))
            }
        }
    }
}

func fact(n int) int {
    if n == 1 {
        return n
    }
    return n * fact(n-1)
}

//func main() {
//    fmt.Println("Fact of locations:", fact(len(locations)))
//    permute(locations)
//    for _, v := range permutations {
//        fmt.Println(v)
//    }
//}
