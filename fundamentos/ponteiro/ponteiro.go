package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiro
	var p *int = nil

	p = &i //pegando o endereço da variavel
	*p++   //desreferenciando
	i++

	fmt.Println(p, *p, &i)

}
