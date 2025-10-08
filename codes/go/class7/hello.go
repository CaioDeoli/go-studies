package main

import "fmt"

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

type ClienteInternacional struct {
	Cliente
	Pais string
}

func main() {
	cliente1 := Cliente{
		"Caio",
		"c@g.com",
		9876,
	}

	cliente2 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Andréa",
			Email: "a@g.com",
			CPF:   4585,
		},
		Pais: "Brasil",
	}

	fmt.Printf("Nome: %s. Email: %s, CPF: %d\n", cliente1.Nome, cliente1.Email, cliente1.CPF)

	fmt.Printf("Nome: %s. Email: %s, Pais: %s\n", cliente2.Nome, cliente2.Email, cliente2.Pais)
}
