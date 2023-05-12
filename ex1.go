//Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
//A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.

package main

import "fmt"

func soma(ptr *int, x int) {
	*ptr = 0
	i := 0
	for i = 0; i < x; i++ {
		*ptr = *ptr + i
	}
	fmt.Println(*ptr)
}

func main() {
	y := 0
	z := 3
	soma(&y, z)

}
