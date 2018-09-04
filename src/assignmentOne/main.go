package main

import (
	"fmt"
)

func main() {
	n := []int{0, 2, 4, 5, 7, 5, 6, 8, 7, 10}

	for _, cn := range n {
		if cn%2 == 0 {
			fmt.Println(cn, "is even")
		} else {
			fmt.Println(cn, "is odd")
		}
	}
}
