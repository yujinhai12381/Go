/*
* File Name: equal.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 一 11/20 10:42:15 2017
*/
package main

import "fmt"


type foo struct {
    a, b, c int
}

func main() {
    s1 := foo {1, 2, 3}
    s2 := foo {1, 2, 3}

    if s1 == s2 {
        fmt.Println("==========")
    }
}
