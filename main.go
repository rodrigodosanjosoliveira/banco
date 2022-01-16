package main

import (
	"fmt"
	"github.com/rodrigodosanjosoliveira/banco/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)

	fmt.Println(contaDoDenis.ObterSaldo())
}
