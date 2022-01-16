package contas

import (
	"github.com/rodrigodosanjosoliveira/banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if !podeSacar {
		return "Saldo insuficiente"
	} else {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Valor do depósito menor que zero", c.saldo
	} else {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia >= c.saldo || valorDaTransferencia <= 0 {
		return false
	} else {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
