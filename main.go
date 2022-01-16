package main

import (
	"fmt"

	"github.com/rodrigodosanjosoliveira/banco/clientes"
	"github.com/rodrigodosanjosoliveira/banco/contas"
)

func main() {
	contaDoBruno := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bruno",
			CPF:       "123.456.789-00",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 123,
		NumeroConta:   22222,
	}

	fmt.Println(contaDoBruno)
}
