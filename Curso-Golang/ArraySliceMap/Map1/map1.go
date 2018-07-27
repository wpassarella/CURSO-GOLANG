package main

import (
	"fmt"
)

func main() {
	//Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678911] = "Aluno1"
	aprovados[12345678912] = "Aluno2"
	aprovados[12345678913] = "Aluno3"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678912])
	delete(aprovados, 12345678912)
	fmt.Println(aprovados[12345678912])
	fmt.Println(aprovados)
}
