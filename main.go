package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var titular string = "André"
	var numeroAgencia int = 598
	var numeroConta int = 123456
	var saldo float64 = 125.5

	fmt.Println(ContaCorrente{titular: titular, numeroAgencia: numeroAgencia, numeroConta: numeroConta, saldo: saldo})
}
