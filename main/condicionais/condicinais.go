package main

import "fmt"

func main() {

	fmt.Printf("%v é primo = %v\n", 30, ehPrimo(30))
	fmt.Printf("%v é primo = %v\n", 17, ehPrimo(17))
	fmt.Printf("Mês %v é %v\n", 12, meses(12))
	fmt.Printf("Mês %v é %v\n", 13, meses(13))

}

func ehPrimo(numero int) bool {
	count := 0

	for i := 1; i <= numero; i++ {
		if numero % i == 0 {
			count++
		}
	}
	return count == 2
}

func meses(temperatura int) string {
	switch temperatura {
	case 1:
		return "Janeiro"
	case 2:
		return "Fevereiro"
	case 3:
		return "Março"
	case 4:
		return "Abril"
	case 5:
		return "Maio"
	case 6:
		return "Junho"
	case 7:
		return "Julho"
	case 8:
		return "Agosto"
	case 9:
		return "Setembro"
	case 10:
		return "Outubro"
	case 11:
		return "Novembro"
	case 12:
		return "Dezembro"
	default:
		return "Inválido"
	}
}
