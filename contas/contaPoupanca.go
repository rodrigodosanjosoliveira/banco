package contas

import "github.com/rodrigodosanjosoliveira/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if !podeSacar {
		return "Saldo insuficiente"
	} else {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Valor do depósito menor que zero", c.saldo
	} else {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
