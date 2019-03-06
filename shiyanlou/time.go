/*
* File Name: time.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三  3/ 6 16:17:40 2019
*/
package main

import "time"
import "fmt"

func main() {
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    for {
        select {
        case res := <-c1:
            fmt.Println(res)
        case <-time.After(time.Second * 1):
            fmt.Println("timeout 1")
        }
    }
}
