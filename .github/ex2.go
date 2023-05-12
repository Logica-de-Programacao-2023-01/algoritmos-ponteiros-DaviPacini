//Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro
//é par ou ímpar. A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar

package main

import "fmt"

func num(ptr *int) {
	div := *ptr % 2
	if div == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
	fmt.Println(*ptr)
}

func main() {
	var x int = 3
	num(&x)
}
