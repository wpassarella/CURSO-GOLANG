package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeatual int
}

type ferrari struct {
	carro       //campos anonimos
	turboligado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeatual = 0
	f.turboligado = true

	fmt.Printf("A Ferrari %s está com o Turbo ligado? %v\n", f.nome, f.turboligado)
}
