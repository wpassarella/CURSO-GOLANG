package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	//adiciona dados ao Slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	//adiciona dados acima da capacidade do array
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
