package main

import "fmt"

type Pessoa struct {
	nome string
	idade int
}

func main() {
	ponteiroBasico()
	ponteiroComStruct()
}

func ponteiroBasico() {
	numero := 42

	var ponteiro = &numero

	fmt.Printf("O valor do ponteiro é %v \n", *ponteiro)

	*ponteiro = 80

	fmt.Printf("O novo valor do ponteiro é %v \n", numero)
}

func ponteiroComStruct() {
	pessoa := Pessoa {"Joshua", 44}

	ponteiro := &pessoa

	fmt.Printf("O valor do ponteiro é %v \n", *ponteiro)

	ponteiro.nome = "Joshua Maia"

	fmt.Printf("O novo valor do ponteiro é %v \n", pessoa)
}