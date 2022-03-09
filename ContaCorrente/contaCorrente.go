package contas

import cli "banco/Clientes"

type ContaCorrente struct {
	Titular cli.Titular
	Agencia int64
	Conta   int64
	Saldo   float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor < 0 {
		return "O valor de saque precisa ser maior do que 0"
	}
	if c.Saldo < valor {
		return "Saldo insuficiente"
	} else {
		c.Saldo -= valor
		return "Saque realizado com sucesso"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor < 0 {
		return "O valor de saque precisa ser maior do que 0", 0
	} else {
		c.Saldo += valor
		return "Depósito realizado com sucesso, o novo saldo é de:", c.Saldo
	}
}
