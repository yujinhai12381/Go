/*
* File Name: main.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 一  8/28 14:02:05 2017
*/
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var ret int
    x := 1534236469 
    ret = reverse(x)
    fmt.Println(x, ret)
    y := 123
    ret = reverse(y)
    fmt.Println(y, ret)
    z := -2147483648
    retz := reverse(z)
    fmt.Println(z, retz)
    n := -123
    retn := reverse(n)
    fmt.Println(n, retn)
    m := 1563847412
    retm := reverse(m)
    fmt.Println(m, retm)
}

func  reverse(x int) int {
    s := fmt.Sprintf("%d", x)
    runes := []rune(s)
    start := 0

    if ('-' == runes[0]) {
        start = 1
    }

    for from, to := start, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    
    str := string(runes)
    ret, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("字符串出错")
    }

    max := (1<<31)
    min := 0-((1<<31))

    if ret>max || ret<min {
        ret = 0
    }

    return ret
}
