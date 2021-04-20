package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaGeovana := ContaCorrente{titular: "Geovana",
		saldo: 125.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDaGeovana)
	fmt.Println(contaDaBruna)

}
