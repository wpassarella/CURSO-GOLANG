package main

import (
	"fmt"
)

func main() {
	funcsESlarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESlarios["Rafael Filho"] = 1350.0
	delete(funcsESlarios, "Inexistente")

	for nome, salario := range funcsESlarios {
		fmt.Println(nome, salario)
	}

}
