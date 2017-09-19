/*
* File Name: for_2.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 二  9/19 11:28:14 2017
*/
package main

import "fmt"

func main() {
    var arr [10]int

    for i := 0; i < 10; i++ {
        arr[i] = i
    }

    fmt.Printf("%v\n", arr)
}
