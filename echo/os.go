/*
* File Name: os.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 日  8/ 6 10:17:46 2017
*/
package main

import (
    "fmt"
    "os"
)

func main () {
    var s, seq string

    for i := 0; i < len(os.Args); i++ {
        s += seq + os.Args[i]
        seq = " "
    }

    fmt.Println(s)
}


