//Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.

package main

import "fmt"

func main() {
	var s string = "rabo"
	var ptr *string = &s

	x := len(*ptr)

	for x = len(*ptr) - 1; x >= 0; x-- {
		fmt.Print(string(s[x]))
	}

}
