/*
* File Name: type_t.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三  3/ 6 21:22:31 2019
*/
package main

import (
    "fmt"
)

func main() {
    x := uint16(65000)
    y := int16(x)

    fmt.Printf("type and value of x is: %T and %d\n", x, x)
    fmt.Printf("type and value of y is: %T and %d\n", y, y)

    var i interface{} = 99
    var s interface{} = []string{"left", "right"}
    j := i.(int)
    fmt.Printf("type and value of j is: %T and %d\n", j, j)

    if s, ok := s.([]string); ok {
        fmt.Printf("%T -> %q\n", s, s)
    }
}
