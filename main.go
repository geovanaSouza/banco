package main

import (
	"fmt"

	"github.com/geovanaSouza/banco/contas"
)

func main() {

	fmt.Println("COMPARAÇÃO DE STRUCTS")
	contaDaGeovana := contas.ContaCorrente{Titular: "Geovana",
		Saldo: 125.5}
	contaDaGeovana2 := contas.ContaCorrente{Titular: "Geovana",
		Saldo: 125.5}

	fmt.Println("contaDaGeovana:", contaDaGeovana)
	fmt.Println("contaDaGeovana2", contaDaGeovana2)
	fmt.Println("Comparação da contaDaGeovana e contaDaGeovana2", contaDaGeovana == contaDaGeovana2)

	contaDaGeovana2.NumeroAgencia = 580
	fmt.Println("Nova agência da conta Geovana2:", contaDaGeovana2.NumeroAgencia)
	fmt.Println("Nova comparação da contaDaGeovana e contaDaGeovana2:", contaDaGeovana == contaDaGeovana2)
	fmt.Println()

	fmt.Println("NOVA CONTA CORRENTE")
	contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println("contaDaBruna", contaDaBruna)
	fmt.Println()

	fmt.Println("PONTEIROS E STRUCT")
	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

	fmt.Println("Conteúdo ponteiro contaDaCris:", contaDaCris)
	fmt.Println("Conteúdo do endereço de memória para o qual o ponteiro direciona (*contaDaCris):", *contaDaCris)
	fmt.Println()

	var contaDaCris2 *contas.ContaCorrente
	fmt.Println("Conteúdo ponteiro contaDaCris2 antes da criação de nova ContaCorrente", contaDaCris2)
	fmt.Println("Endereço de memória contaDaCris2", &contaDaCris2)
	fmt.Println()

	contaDaCris2 = new(contas.ContaCorrente)
	fmt.Println("Conteúdo ponteiro contaDaCris2 depois da criação de nova ContaCorrente", contaDaCris2)
	fmt.Println("Conteúdo do endereço de memória para o qual o ponteiro direciona (*contaDaCris2):", *contaDaCris2)
	fmt.Println("Endereço de memória", &contaDaCris2)
	fmt.Println()

	contaDaCris2.Titular = "Cris"
	contaDaCris2.Saldo = 500

	fmt.Println("ponteiro contaDaCris:", contaDaCris)
	fmt.Println("ponteiro contaDaCris2:", contaDaCris2)
	fmt.Println("Comparando ponteiros contaDaCris e contaDaCris2:", contaDaCris == contaDaCris2)
	fmt.Println("Comparando conteúdos dos endereços de memória para o qual os ponteiros direcionam (*contaDaCris e *contaDaCris2):", *contaDaCris == *contaDaCris2)
	fmt.Println()

	fmt.Println("FUNÇÃO SAQUE")
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println("Saque na conta da", contaDaSilvia.Titular)
	valorDoSaque := 300.
	fmt.Println(contaDaSilvia.Sacar(valorDoSaque))

	fmt.Println(contaDaSilvia.Sacar(600))
	fmt.Println(contaDaSilvia.Sacar(-100))
	fmt.Println(contaDaSilvia.Sacar(200))

	status, valor := contaDaSilvia.Depositar(500)
	fmt.Println(status, "\nNovo saldo da conta:", valor)
	fmt.Println()

	status, valor = contaDaSilvia.Depositar(500)
	fmt.Println(status, "\nNovo saldo da conta:", valor)
	fmt.Println()

	status, valor = contaDaSilvia.Depositar(-1000)
	fmt.Println(status, "\nNovo saldo da conta:", valor)
	fmt.Println()

	fmt.Println("FUNÇÃO TRANSFERÊNCIA")
	contaDoGustavo := new(contas.ContaCorrente)
	contaDoGustavo.Titular = "Gustavo"
	contaDoGustavo.Saldo = 100

	fmt.Println("contaDaSilvia:", contaDaSilvia)
	fmt.Println("contaDoGustavo:", *contaDoGustavo)
	fmt.Println()

	statusTransferencia := contaDaSilvia.Transferir(200, contaDoGustavo)
	fmt.Println("Status da Transferencia:", statusTransferencia)
	fmt.Println()

	statusTransferencia = contaDoGustavo.Transferir(400, &contaDaSilvia)
	fmt.Println("Status da Transferência:", statusTransferencia)
	fmt.Println()

	statusTransferencia = contaDoGustavo.Transferir(-200, &contaDaSilvia)
	fmt.Println("Status da Transferência:", statusTransferencia)
	fmt.Println()

}
