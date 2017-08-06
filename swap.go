/*
* File Name: swap.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三  8/ 2 17:33:05 2017
*/
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

