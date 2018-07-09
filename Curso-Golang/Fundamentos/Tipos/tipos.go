package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Números interiors
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (só postivios)... uint uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Ola me nome é Wagner"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Wagner!`
	fmt.Println("O tamanho da string é", len(s2))
}
