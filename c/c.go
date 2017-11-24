/*
* File Name: c.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 五 11/24 16:28:49 2017
*/
package main

/*
#include <stdio.h>

void hello() {
    printf("hello, c go!!  I am Haha!\n");
}
*/
import "C"

func Hello() {
    C.hello()
}

func main() {
    Hello()
}
