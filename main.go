package main

import "fmt"

func main() {
	var titular string = "André"
	var numeroAgencia int = 598
	var numeroConta int = 123456
	var saldo float64 = 125.5

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	var titular2 string = "Fulano"
	var numeroAgencia2 int = 122
	var numeroConta2 int = 556633
	var saldo2 float64 = 999.3

	fmt.Println(titular2, numeroAgencia2, numeroConta2, saldo2)
}
