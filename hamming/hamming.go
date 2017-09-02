/*
* File Name: hamming.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 六  9/ 2 17:24:30 2017
*/
package main

import "fmt"

func main() {
    ret := hammingDistance(1,4)
    fmt.Printf("%d\n", ret)
}

func hammingDistance(x int, y int) int {
    hamming := x ^ y
    ret := 0

    for i := 0; i<32; i++ {
        num := hamming&0x1
        ret += num
        hamming = hamming>>1
    }

    return ret
}

