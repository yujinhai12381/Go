/*
* File Name: main.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 日  8/ 6 10:52:43 2017
*/
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Println(os.Args[1:])
}
