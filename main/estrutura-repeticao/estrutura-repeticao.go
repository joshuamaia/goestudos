package main

import (
	"fmt"
) 


func main() {
  estruturaFor()
	estruturaForComoWhile()
	estruturaForComString()
}

func estruturaForComString() {
	s := "Estamos aqui"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
}

func estruturaFor(){
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %v", i)
	}
}

func estruturaForComoWhile() {
	sum := 0
	for  sum < 100 {
		sum = sum + 8
		fmt.Println("Sum =", sum)
	}
}