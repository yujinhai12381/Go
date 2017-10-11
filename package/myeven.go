/*
* File Name: myeven.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三 10/11 12:12:19 2017
*/
package main

import (
    "even"
    "fmt"
)

func main() {
    i := 6
    fmt.Printf("Is %d even ? %v \n", i, even.Even(i))
}
