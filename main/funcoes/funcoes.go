package main

import (
	"fmt"
) 


func main() {
  fmt.Println("Soma =", soma(4,5))
	fmt.Println("Divisão =", divisao(1,0))
	var valor, somatorio = valorJuros(3, 1000, 0.1)
	fmt.Printf("O valor ficou %v e o total %v", valor, somatorio)
}

func soma(a int, b int) int {
	return a + b
}

func divisao(a float64, b float64) float64 {
	resultado := a/b
	return resultado
}

func valorJuros(meses int, valor float64, juros float64) (float64, float64) {
	soma := 0.0
	for i := 1; i <= meses; i++ {
		valor = valor + valor * juros
		soma = soma + valor
	}
	return valor, soma
}