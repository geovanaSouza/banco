package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	fmt.Println("Saldo da conta:", c.Saldo)
	fmt.Println("Sacando", valorDoSaque, "da conta")
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		fmt.Println("Novo valor de saldo da conta:", c.Saldo)
		return "Saque relizado com sucesso\n"
	} else if valorDoSaque < 0 {
		return "Não é possível fazer saque negativo\n"
	} else {
		return "Saldo insuficiente\n"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	fmt.Println("Saldo da conta:", c.Saldo)
	fmt.Println("Depositando", valorDoDeposito, "na conta")
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do deposito menor que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	fmt.Println("Saldo da conta Origem:", c.Saldo)
	fmt.Println("Saldo da conta Destino:", contaDestino.Saldo)
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		fmt.Println("Novo saldo da conta Origem:", c.Saldo)
		fmt.Println("Novo saldo da conta Destino:", contaDestino.Saldo)
		return true
	} else {
		return false
	}
}
