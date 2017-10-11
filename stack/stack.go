/*
* File Name: stack.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三 10/11 10:44:53 2017
*/
package main

import "fmt"

type stack struct {
    i int
    data [10]int
}

func main() {
    var s stack
    s.push(25)
    s.push(14)
    fmt.Printf("stack : %v\n", s)
    //a := s.pop()
    //fmt.Printf("a : %v\n", a)
    s.pop()
    fmt.Printf("stack : %v\n", s)
}

func (s *stack) push(k int) {
    s.data[s.i] = k;
    s.i++
}

func (s *stack) pop() int {
    s.i--
    ret := s.data[s.i]
    s.data[s.i] = 0

    return ret
}

