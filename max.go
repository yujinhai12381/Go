/*
* File Name: max.go
* Desc:
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 三  8/ 2 17:18:34 2017
*/
package main

import "fmt"

func main() {
    var a int = 100
    var b int = 200
    var ret int

    ret = max(a, b);

    fmt.Printf( "max is :%d \n", ret )
}

func max(num1, num2 int) int {
    var ret int

    if (num1 > num2) {
        ret = num1
    } else {
        ret = num2
    }

    return ret
}

