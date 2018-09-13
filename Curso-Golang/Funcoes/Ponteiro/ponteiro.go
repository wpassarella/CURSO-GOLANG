package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n) //ponteiro por valor
	fmt.Println(n)

	inc2(&n) //ponteiro por referência
	fmt.Println(n)

}
