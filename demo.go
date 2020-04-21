/*
* File Name: demo.go
* Desc:
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 2020年04月19日 星期日 11时13分14秒
 */
package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
