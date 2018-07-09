package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aluno Aprovado com nota: ", nota)
	} else {
		fmt.Println("Aluno Reprovado com nota: ", nota)
	}
}
func main() {
	imprimirResultado(7.9)
	imprimirResultado(3.0)
}
