package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//forma reduzida
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é ", area)
}
