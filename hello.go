/*
* File Name: hello.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三  2/27 15:10:13 2019
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    target := "world"

    if len(os.Args) > 1 {
        target = os.Args[1]
    }

    fmt.Println("Hello", target)
}
