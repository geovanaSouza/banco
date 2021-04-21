package contas

import (
	"fmt"

	"github.com/geovanaSouza/banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	fmt.Println("saldo da conta:", c.saldo)
	fmt.Println("Sacando", valorDoSaque, "da conta")
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		fmt.Println("Novo valor de saldo da conta:", c.saldo)
		return "Saque relizado com sucesso\n"
	} else if valorDoSaque < 0 {
		return "Não é possível fazer saque negativo\n"
	} else {
		return "saldo insuficiente\n"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	fmt.Println("saldo da conta:", c.saldo)
	fmt.Println("Depositando", valorDoDeposito, "na conta")
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	fmt.Println("saldo da conta Origem:", c.saldo)
	fmt.Println("saldo da conta Destino:", contaDestino.saldo)
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		fmt.Println("Novo saldo da conta Origem:", c.saldo)
		fmt.Println("Novo saldo da conta Destino:", contaDestino.saldo)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
