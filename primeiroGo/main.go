package main

import (
	"fmt"

	"github.com/PedroHODL/bootcamp/primeiroGo/calc"
)

func main() {
	fmt.Println("Olá Mundo!")

	var (
		nome  = "Pedro"
		idade = 20
	)
	var hetero bool = true

	fmt.Printf("Meu nome é: %s, Idade: %d, opção sexual: %t\n", nome, idade, hetero)

	idade = calc.Soma(idade, 20)

	fmt.Println(idade)
}
