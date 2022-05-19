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
	var homem bool = true
	fmt.Printf("Meu nome é: %s, Idade: %d, é homem? %t\n", nome, idade, homem)

	idade = calc.Soma(idade, 20)
	fmt.Println(idade)
}
