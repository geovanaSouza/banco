package main

import (
	"fmt"

	"github.com/geovanaSouza/banco/clientes"
	"github.com/geovanaSouza/banco/contas"
)

func main() {

	fmt.Println("COMPARAÇÃO DE STRUCTS")
	contaDaGeovana := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Geovana"}}
	contaDaGeovana.Depositar(125.5)
	contaDaGeovana2 := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Geovana"}}
	contaDaGeovana2.Depositar(125.5)

	fmt.Println("contaDaGeovana:", contaDaGeovana)
	fmt.Println("contaDaGeovana2", contaDaGeovana2)
	fmt.Println("Comparação da contaDaGeovana e contaDaGeovana2", contaDaGeovana == contaDaGeovana2)

	contaDaGeovana2.NumeroAgencia = 580
	fmt.Println("Nova agência da conta Geovana2:", contaDaGeovana2.NumeroAgencia)
	fmt.Println("Nova comparação da contaDaGeovana e contaDaGeovana2:", contaDaGeovana == contaDaGeovana2)
	fmt.Println()

	fmt.Println("NOVA CONTA CORRENTE")
	contaDaBruna := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Bruna", CPF: "110.677.330-52", Profissao: "SRE"}, NumeroAgencia: 222, NumeroConta: 111222}
	contaDaBruna.Depositar(200)
	fmt.Println("contaDaBruna", contaDaBruna)
	fmt.Println()

	fmt.Println("PONTEIROS E STRUCT")
	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular.Nome = "Cris"
	contaDaCris.Depositar(500)

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

	contaDaCris2.Titular.Nome = "Cris"
	contaDaCris2.Depositar(500)

	fmt.Println("ponteiro contaDaCris:", contaDaCris)
	fmt.Println("ponteiro contaDaCris2:", contaDaCris2)
	fmt.Println("Comparando ponteiros contaDaCris e contaDaCris2:", contaDaCris == contaDaCris2)
	fmt.Println("Comparando conteúdos dos endereços de memória para o qual os ponteiros direcionam (*contaDaCris e *contaDaCris2):", *contaDaCris == *contaDaCris2)
	fmt.Println()

	fmt.Println("FUNÇÃO SAQUE")
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular.Nome = "Silvia"
	contaDaSilvia.Depositar(500)

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
	contaDoGustavo.Titular.Nome = "Gustavo"
	contaDoGustavo.Depositar(100)

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

	fmt.Println("CRIANDO CONTA E TITULAR JUNTO usando composição")
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456}
	contaDoBruno.Depositar(100)
	fmt.Println(contaDoBruno)
	fmt.Println()

	fmt.Println("CRIANDO CLIENTE PRIMEIRO, DEPOIS CONTA usando composição")
	clienteBruno2 := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor Go"}
	contaDoBruno2 := contas.ContaCorrente{Titular: clienteBruno2, NumeroAgencia: 123, NumeroConta: 123456}
	contaDoBruno2.Depositar(100)
	fmt.Println(contaDoBruno2)
	fmt.Println()

	fmt.Println("LIDANDO COM VISIBILIDADE DO SALDO")
	contaExemplo := new(contas.ContaCorrente)
	contaExemplo.Depositar(100)
	// Forçando erro após modificar saldo para privado
	// Error: contaExemplo.Saldo undefined (type *contas.ContaCorrente has no field or method Saldo)
	// contaExemplo.Saldo = -100
	fmt.Println(contaExemplo.ObterSaldo())
}
