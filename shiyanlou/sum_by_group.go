/*
* File Name: sum_by_group.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三  2/27 15:28:09 2019
*/
package main

import "fmt"

func sum(a []int, result chan int) {
    sum :=0
    for _, v := range a {
    //    fmt.Println(v)
        sum += v
    }

    result <-sum
}

func main() {
    a := []int{2, 3, 5, 6, 10, -5, 1, 0}
    result := make(chan int)
    go sum(a[:len(a)/2], result)
    go sum(a[len(a)/2:], result)
    x, y := <-result, <-result

    fmt.Println(x, y, x + y)
    //x := <-result
    //fmt.Println(x)
}
