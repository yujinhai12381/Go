/*
* File Name: main.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 日  8/ 6 10:35:59 2017
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    s, seq := "", ""

    for _, arg := range os.Args[1:] {
        s += seq + arg
        seq = " "
    }

    fmt.Println(s)
}
