package main

import "fmt"

type ContaCorrente struct {
	titular 		string
	numeroAgencia	int
	numeroConta 	int
	saldo 			float64
}

func main() {
	contaDoHenrique := ContaCorrente{titular: "Henrique", numeroAgencia: 589, numeroConta: 123456, saldo: 300.0}
	fmt.Println(contaDoHenrique)
}