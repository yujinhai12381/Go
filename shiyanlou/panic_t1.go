/*
* File Name: panic_t1.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 四  3/ 7 17:50:48 2019
*/
package main

import (
    "fmt"
    "math"
)

func ContvertIntToInt16(x int) int16 {
    if math.MinInt16 <= x && x <= math.MaxInt16 {
        return int16(x)
    }

    panic(fmt.Sprintf("%d is out of int16 range", x)) // 手动触发panic
}

func Int16FromInt(x int) (i int16, err error) {
    defer func() {   // 延迟执行匿名函数，并使用recover捕捉了panic，并将panic转换为了error
        if e := recover(); e != nil {
            i = 11
            err = fmt.Errorf("%v", e)
        }
    }()
    i = ContvertIntToInt16(x)
    return i, nil
}

func main() {
    if i, e := Int16FromInt(655567); e != nil {
        fmt.Printf("%d\n", i)
        fmt.Printf("%v\n", e)
    } else {
        fmt.Printf("no errors\n")
    }
}

