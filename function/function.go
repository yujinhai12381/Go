/*
* File Name: function.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三 10/11 10:16:09 2017
*/
package main

import "fmt"

func main() {
    p2 := plusTwo()
    fmt.Printf("%v\n", p2(2))
}

func plusTwo() func(int) int {
    return func(x int) int { return x + 2 }
}

