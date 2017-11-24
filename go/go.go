/*
* File Name: go.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 五 11/24 15:58:06 2017
*/
package main

import "fmt"
import "time"

func Add(x, y int) {
    z := x + y
    fmt.Println(z)
}

func main() {
    for i := 0; i < 10; i++ {
        go Add(i, i)
    }
    time.Sleep(3*time.Second)
}
