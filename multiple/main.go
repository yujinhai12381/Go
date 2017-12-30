package main

import "fmt"

func vals(a, b int) (int, int) {
	return a, b
}

func main() {
	a, b := vals(3, 7)
    fmt.Println(a)
	fmt.Println(b)
	
	_, c := vals(6, 8)
    fmt.Println(c)
}
