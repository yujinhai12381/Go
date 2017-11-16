/*
* File Name: channel.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 四 11/16 19:06:45 2017
*/
package main

import "fmt"

func main() {
    ch := make(chan int)
    go shower(ch)

    ch <- 666 

    for i:= 0; i< 10; i++ {
        ch <- i
    }
}

func shower(c chan int) {
    for {
        j := <- c
        fmt.Printf("%d\n", j)
    }
}
